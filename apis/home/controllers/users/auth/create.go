package auth

import (
	"basic/apis/common/validator"
	dtousers "basic/apis/home/dto/users"
	"basic/apis/home/service/users"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Create
// @Tags User
// @Summary	新增认证
// @Param		realName	body		string	true	"真实姓名"
// @Param		number		body		string	true	"卡号"
// @Param		photo1		body		string	true	"证件照1"
// @Param		photo2		body		string	true	"证件照2"
// @Param		telephone	body		string	true	"电话"
// @Param		address		body		string	true	"地址"
// @Success	200	{string}	string	"ok"
// @Router		/api/v1/user/auth/create [post]
func Create(c *fiber.Ctx) error {
	params := new(dtousers.AuthCreateParams)
	err := c.BodyParser(&params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	验证参数
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	调用服务层注册流程
	adminId, userId := utils.GetContextClaims(c)
	lang := c.Get("Accept-Language")
	data, err := users.AuthCreate(adminId, userId, lang, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
