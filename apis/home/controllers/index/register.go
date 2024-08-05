package index

import (
	"basic/apis/common/validator"
	dtousers "basic/apis/home/dto/users"
	"basic/apis/home/service/users"
	"basic/utils"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

// Register 用户注册
// @Tags Basic
// @Summary	用户注册
// @Param		username	body		string					true	"用户名"
// @Param		password	body		string					true	"密码"
// @Param		captchaId	body		string					false	"验证码ID"
// @Param		captchaVal	body		string					false	"验证码值"
// @Param		email		body		string					false	"邮箱"
// @Param		telephone	body		string					false	"手机号码"
// @Param		securityKey	body		string					false	"安全密钥"
// @Param		code		body		string					false	"用户邀请码"
// @Success	200			{object}	dtousers.UserLoginData
// @Router		/api/v1/register [post]
func Register(c *fiber.Ctx) error {
	params := &dtousers.UserRegisterParams{}
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
	data, err := users.UserRegister(originParseUrl.Host, lang, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	返回数据
	return c.JSON(utils.SuccessJson(data))
}
