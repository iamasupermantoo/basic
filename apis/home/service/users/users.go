package users

import (
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/apis/common/service/users"
	dtousers "basic/apis/home/dto/users"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"
	"errors"
	"time"

	"github.com/dchest/captcha"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

// UserLogin 用户登录
func UserLogin(host, lang string, params *dtousers.UserLoginParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	域名获取管理ID
	adminId := rds.RedisFindDomainToAdminId(rdsConn, host)
	settingAdminId := rds.RedisFindAdminIdToSettingAdminId(rdsConn, adminId)

	//	获取登录配置
	loginSettingList := utils.StringToBoolMaps(rds.RedisFindAdminSettingField(rdsConn, adminId, "templateLogin"))
	if _, ok := loginSettingList["showVerify"]; ok {
		if loginSettingList["showVerify"] && !captcha.VerifyString(params.CaptchaId, params.CaptchaVal) {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "incorrectCode"))
		}
	}

	//	检测管理是否存在， 并且密码是否匹配
	userInfo := &dtousers.UserInfo{}
	if result := database.Db.Model(&models.User{}).
		Where("username = ?", params.UserName).
		Where("status = ?", models.AdminUserStatusActive).
		Find(userInfo); result.Error != nil || userInfo.Password != utils.PasswordEncrypt(params.Password) || userInfo.AdminId != settingAdminId {
		return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "accountOrPasswordIncorrect"))
	}

	//	生成Token
	userLoginData := &dtousers.UserLoginData{}
	userLoginData.Token = models.NewServiceHomeToken(userInfo.AdminId, userInfo.Id)
	userLoginData.UserInfo, _ = UserInfo(userInfo.Id)
	return userLoginData, nil
}

// UserRegister 用户注册
func UserRegister(host, lang string, params *dtousers.UserRegisterParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	域名获取管理ID
	adminId := rds.RedisFindDomainToAdminId(rdsConn, host)

	//	获取注册配置文件
	registerSettingList := utils.StringToBoolMaps(rds.RedisFindAdminSettingField(rdsConn, adminId, "templateRegister"))
	//	验证码必须必须
	if _, ok := registerSettingList["showVerify"]; ok {
		if registerSettingList["showVerify"] && !captcha.VerifyString(params.CaptchaId, params.CaptchaVal) {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "incorrectCode"))
		}
	}

	//	是否关闭注册
	if _, ok := registerSettingList["closeRegister"]; ok {
		if registerSettingList["closeRegister"] {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "notRegister"))
		}
	}

	//	安全密钥是否必须
	userSecurityKey := ""
	if _, ok := registerSettingList["showSecurityPass"]; ok {
		if registerSettingList["showSecurityPass"] && params.SecurityKey == "" {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "securityBeEmpty"))
		}
		userSecurityKey = utils.PasswordEncrypt(params.SecurityKey)
	}

	//	手机号码是否为空
	userTelephone := ""
	if _, ok := registerSettingList["showTelephone"]; ok {
		if registerSettingList["showTelephone"] && params.Telephone == "" {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "telephoneBeEmpty"))
		}
		userTelephone = params.Telephone
	}

	//	邮箱不能为空
	userEmail := ""
	if _, ok := registerSettingList["showEmail"]; ok {
		if registerSettingList["showEmail"] && params.Email == "" {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "emailBeEmpty"))
		}
		userEmail = params.Email
	}

	//	是否需要邀请码
	parentId := 0
	if _, ok := registerSettingList["isInvite"]; ok {
		if registerSettingList["isInvite"] && params.Code == "" {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "inviteBeEmpty"))
		}
	}
	if params.Code != "" {
		inviteInfo, err := rds.RedisFindInviteCodeInfo(rdsConn, params.Code)
		if err != nil {
			return nil, err
		}

		//	是否覆盖上级ID， 管理ID
		parentId = inviteInfo.UserId
		adminId = inviteInfo.AdminId
	}

	//	插入用户数据
	userInfo := &models.User{
		ParentId:    parentId,
		AdminId:     adminId,
		UserName:    params.UserName,
		Password:    utils.PasswordEncrypt(params.Password),
		SecurityKey: userSecurityKey,
		Telephone:   userTelephone,
		Email:       userEmail,
	}

	registerAward := &dtousers.UserRegisterAward{}
	_ = json.Unmarshal([]byte(rds.RedisFindAdminSettingField(rdsConn, adminId, "registerAward")), &registerAward)

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		//	判断是否插入失败
		if result := tx.Create(userInfo); result.Error != nil {
			return result.Error
		}

		//	是否有邀请奖励
		if registerAward.Register > 0 {
			_ = users.UserDeposit(tx, &dto.WalletOrderAgreeParams{
				AdminId:  userInfo.AdminId,
				ParentId: userInfo.ParentId,
				UserId:   userInfo.Id,
				SourceId: userInfo.Id,
				BillType: models.WalletBillTypeRegisterAward,
				Money:    registerAward.Register,
			})
		}

		//	是否有上级ID
		if registerAward.Share > 0 && userInfo.ParentId > 0 {
			userParentInfo := &models.User{}
			if result := database.Db.Model(userParentInfo).Where("id = ?", userInfo.ParentId).Take(userParentInfo); result.Error != nil {
				return result.Error
			}

			//	邀请奖励
			_ = users.UserDeposit(tx, &dto.WalletOrderAgreeParams{
				AdminId:  userParentInfo.AdminId,
				ParentId: userParentInfo.ParentId,
				UserId:   userParentInfo.Id,
				SourceId: userInfo.Id,
				BillType: models.WalletBillTypeShareAward,
				Balance:  userParentInfo.Money,
				Money:    registerAward.Register,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	//	生成Token
	userLoginData := &dtousers.UserLoginData{}
	userLoginData.Token = models.NewServiceHomeToken(userInfo.AdminId, userInfo.Id)
	userLoginData.UserInfo, _ = UserInfo(userInfo.Id)
	return userLoginData, nil
}

// UserUpdate 更新当前用户信息
func UserUpdate(adminId int, userId int, params *dtousers.UserUpdateParams) (interface{}, error) {
	if params.BirthdayStr != "" {
		birthdayTime, err := time.Parse("2006/01/02", params.BirthdayStr)
		if err != nil {
			return nil, err
		}
		params.Birthday = int(birthdayTime.Unix())
	}
	result := database.Db.Model(models.User{}).Where("id=?", userId).Updates(&params)
	return "ok", result.Error
}

// UserInfo 获取当前用户信息
func UserInfo(userId int) (interface{}, error) {
	userInfo := &dtousers.HomeUserInfo{}
	result := database.Db.Table("user as u").
		Select("u.id", "u.username", "u.nickname", "u.avatar", "u.email", "u.telephone", "u.sex", "u.money", "u.score", "u.birthday", "u.desc", "u.created_at", "l.level ", "rna.status as authStatus").
		Joins("LEFT JOIN level_order as lo on u.id = lo.user_id AND lo.status = ?", models.UserLevelOrderStatusActive).
		Joins("LEFT JOIN level as l on lo.level_id = l.id").
		Joins("LEFT JOIN real_name_auth as rna on u.id = rna.user_id").
		Where("u.id=?", userId).Take(userInfo)
	return userInfo, result.Error
}

// UserUpdatePassword 更新当前用户密码||安全码
func UserUpdatePassword(adminId int, userId int, lang string, params *dtousers.UserUpdatePasswordParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//获取当前用户信息
	userInfo := &dtousers.UserInfo{}
	result := database.Db.Model(&models.User{}).Where("id=?", userId).Take(userInfo)
	if result.Error == nil {
		switch params.Type {
		//修改密码
		case dtousers.UpdatePassword:
			//输入旧密码错误
			if userInfo.Password != utils.PasswordEncrypt(params.OldPassword) {
				return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "TheOldPasswordIsIncorrect"))
			}
			if result = database.Db.Model(&models.User{}).
				Where("id = ?", userId).
				Update("password", utils.PasswordEncrypt(params.NewPassword)); result.Error != nil {
				return nil, result.Error
			}
		//修改安全密码
		case dtousers.UpdateSecurityPassword:
			//输入旧安全密码错误
			if userInfo.SecurityKey != utils.PasswordEncrypt(params.OldPassword) {
				return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "TheOldSecurityKeyIsIncorrect"))
			}
			if result = database.Db.Model(&models.User{}).
				Where("id = ?", userId).
				Update("security_key", utils.PasswordEncrypt(params.NewPassword)); result.Error != nil {
				return nil, result.Error
			}
		}

	}
	return "ok", result.Error
}
