package user

import (
	"basic/apis/common/validator"
	dtousers "basic/apis/home/dto/users"
	"basic/apis/home/service/users"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// UpdatePassword
// @Summary	更新用户密码或安全码
// @Tags		User
// @Param		type		body		int64	true	"密码类型 (1: 更新密码, 2: 更新安全码)"
// @Param		oldPassword	body		string	true	"旧密码 (长度: 5-28)"
// @Param		newPassword	body		string	true	"新密码 (长度: 5-28)"
// @Success	200			{string}	string
// @Router		/api/v1/user/update/password [post]
func UpdatePassword(c *fiber.Ctx) error {
	params := &dtousers.UserUpdatePasswordParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	验证规则
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	调用服务层
	adminId, userId := utils.GetContextClaims(c)
	lang := c.Get("Accept-Language")
	data, err := users.UserUpdatePassword(adminId, userId, lang, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	返回数据
	return c.JSON(utils.SuccessJson(data))
}
