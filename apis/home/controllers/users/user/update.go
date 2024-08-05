package user

import (
	"basic/apis/common/validator"
	dtousers "basic/apis/home/dto/users"
	"basic/apis/home/service/users"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Update
// @Summary	更新用户信息
// @Tags		User
// @Param		avatar		body		string	false	"头像"
// @Param		nickname	body		string	false	"昵称"
// @Param		email		body		string	false	"邮箱"
// @Param		telephone	body		string	false	"电话"
// @Param		sex			body		int		false	"性别"
// @Param		birthday	body		string	false	"生日"
// @Param		data		body		string	false	"数据"
// @Param		desc		body		string	false	"描述"
// @Success	200			{string}	string
// @Router		/api/vi/user/update [post]
func Update(c *fiber.Ctx) error {
	params := &dtousers.UserUpdateParams{}
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
	data, err := users.UserUpdate(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	返回数据
	return c.JSON(utils.SuccessJson(data))
}
