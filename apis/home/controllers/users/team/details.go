package team

import (
	"basic/apis/common/validator"
	dtousers "basic/apis/home/dto/users"
	"basic/apis/home/service/users"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Details
// @Summary	获取用户团队详情数据
// @Tags		User
// @Success	200	{object}	dtousers.TeamDetailsData
// @Router		/api/v1/user/team/details [post]
func Details(c *fiber.Ctx) error {
	params := &dtousers.TeamDetailsParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	验证规则
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	lang := c.Get("Accept-Language")
	data, err := users.TeamDetails(adminId, userId, lang, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	返回数据
	return c.JSON(utils.SuccessJson(data))
}
