package index

import (
	"basic/apis/common/validator"
	dtousers "basic/apis/home/dto/users"
	"basic/apis/home/service/users"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

// Login
// @Summary	用户登录
// @Tags Basic
// @Param		username	body		string					true	"用户名"
// @Param		password	body		string					true	"密码"
// @Param		captchaId	body		string					false	"验证码ID"
// @Param		captchaVal	body		string					false	"验证码值"
// @Success	200			{object}	dtousers.UserLoginData
// @Router		/api/v1/login [post]
func Login(c *fiber.Ctx) error {
	params := &dtousers.UserLoginParams{}
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
	originParseUrl, _ := url.Parse(c.Get("Origin"))
	lang := c.Get("Accept-Language")
	data, err := users.UserLogin(originParseUrl.Host, lang, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	返回数据
	return c.JSON(utils.SuccessJson(data))
}
