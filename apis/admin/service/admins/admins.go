package admins

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"
	"errors"
	"fmt"
	"github.com/dchest/captcha"
	"net"
)

// AdminLogin 管理登录
func AdminLogin(params *dtoadmins.AdminLoginParams) (interface{}, error) {
	//	检查验证码是否正确
	if !captcha.VerifyString(params.CaptchaId, params.CaptchaVal) {
		return nil, errors.New("验证码不正确")
	}

	//	检测管理是否存在， 并且密码是否匹配
	adminInfo := &dtoadmins.AdminInfo{}
	if result := database.Db.Table("admin_user").
		Where("username = ?", params.UserName).
		Where("status = ?", models.AdminUserStatusActive).
		Take(adminInfo); result.Error != nil || adminInfo.Password != utils.PasswordEncrypt(params.Password) {
		return nil, errors.New("账号或密码不正确")
	}

	//	生成Token
	return &dtoadmins.AdminLoginData{Token: models.NewServiceAdminToken(adminInfo.Id)}, nil
}

// AdminInfo 管理信息
func AdminInfo(adminId int) (interface{}, error) {
	//获取当前管理员信息
	var adminInfo dtoadmins.AdminInfo
	if result := database.Db.Model(&models.AdminUser{}).Where("id=?", adminId).Find(&adminInfo); result.Error != nil {
		return nil, result.Error
	}
	return adminInfo, nil
}

// AdminUpdate 更新当前管理员信息
func AdminUpdate(params *dtoadmins.AdminUpdateParams) (interface{}, error) {
	if result := database.Db.Model(&models.AdminUser{}).Where("id = ?", params.Id).Updates(params); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}

// AdminUpdatePassword 更新当前管理员密码
func AdminUpdatePassword(params *dtoadmins.AdminUpdatePasswordParams) (interface{}, error) {
	//获取当前管理员信息
	adminInfo := &dtoadmins.AdminInfo{}
	if result := database.Db.Model(&models.AdminUser{}).Where("id=?", params.Id).Find(adminInfo); result.Error == nil {

		//实例化模型
		switch params.Type {
		case dtoadmins.UpdatePassword:
			if adminInfo.Password != utils.PasswordEncrypt(params.OldPassword) {
				return nil, errors.New("TheOldPasswordIsIncorrect")
			}
			if result = database.Db.Model(&models.AdminUser{}).
				Where("id = ?", adminInfo.Id).
				Update("password", utils.PasswordEncrypt(params.NewPassword)); result.Error != nil {
				return nil, result.Error
			}
		case dtoadmins.UpdateSecurityPassword:
			if adminInfo.SecurityKey != utils.PasswordEncrypt(params.OldPassword) {
				return nil, errors.New("TheOldPasswordIsIncorrect")
			}
			if result = database.Db.Model(&models.AdminUser{}).
				Where("id = ?", adminInfo.Id).
				Update("security_key", utils.PasswordEncrypt(params.NewPassword)); result.Error != nil {
				return nil, result.Error
			}
		}
	}
	return "ok", nil
}

// Audio 音源
func Audio(adminId int) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	// 提示音设置
	audioSettings := utils.StringToBoolMaps(rds.RedisFindAdminSettingField(rdsConn, adminId, "audioSetting"))

	//充值声音
	if audioSettings["deposit"] {
		var depositNums int64
		if result := database.Db.Model(&models.WalletOrder{}).
			Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
			Where("type IN ?", []int{models.WalletOrderTypeDeposit, models.WalletOrderTypeAssetsDeposit}).
			Where("status = ?", models.WalletOrderStatusActive).
			Count(&depositNums); result.Error != nil {
			return nil, result.Error
		}

		if depositNums > 0 {
			return map[string]string{"title": fmt.Sprintf("您有 %v 充值订单未处理", depositNums), "source": "/assets/mp3/deposit.mp3"}, nil
		}
	}

	//提现声音
	if audioSettings["withdraw"] {
		var withdrawNums int64
		if result := database.Db.Model(&models.WalletOrder{}).
			Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
			Where("type IN ?", []int{models.WalletOrderTypeWithdraw, models.WalletOrderTypeAssetsWithdraw}).
			Where("status = ?", models.WalletOrderStatusActive).
			Count(&withdrawNums); result.Error != nil {
			return nil, result.Error
		}
		if withdrawNums > 0 {
			return map[string]string{"title": fmt.Sprintf("您有 %v 提现订单未处理", withdrawNums), "source": "/assets/mp3/withdraw.mp3"}, nil
		}
	}

	return map[string]string{"title": "", "source": ""}, nil
}

// AdminLogsIndex 操作日志列表
func AdminLogsIndex(adminId int, userId int, params *dtoadmins.AdminLogsIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: make([]*dtoadmins.AdminLogsIndexData, 0)}
	var ipInt uint32
	if params.Ip != "" {
		ipInt = utils.Ip4ToInt(net.ParseIP(params.Ip))
	}
	//添加并过滤where条件
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE  ?", "%"+params.AdminName+"%")).
		Like("name LIKE ?", params.Name+"%").
		Like("route LIKE ?", params.Route+"%").
		EqInt("ip = ?", int(ipInt)).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	//设置查询条件
	if result := database.Db.Table("admin_logs as al").
		Select("au.username as adminName,al.id,al.name,INET_NTOA(al.ip) as ip,al.headers,al.route,al.body,al.data,al.created_at").
		Joins("Left JOIN admin_user au ON al.admin_id = au.id").
		Where("au.status=?", models.AdminUserStatusActive).
		Where("al.admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).Count(&data.Count).Scopes(utils.Paginate(params.Pagination)).Scan(&data.Items); result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}
