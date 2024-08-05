package team

import (
	"basic/apis/common/validator"
	dtousers "basic/apis/home/dto/users"
	"basic/apis/home/service/users"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index 团队列表
// @Summary	获取团队成员列表
// @Tags		User
// @Param		id	path		int					true	"用户ID"
// @Success	200	{object}	dtousers.TeamIndexInfo	"团队成员列表"
// @Router		/api/v1/user/team/index [post]
func Index(c *fiber.Ctx) error {
	params := &dtousers.TeamIndexParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	验证规则
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//调用服务层
	adminId, userId := utils.GetContextClaims(c)
	lang := c.Get("Accept-Language")
	data, err := users.TeamIndex(adminId, userId, lang, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	返回数据
	return c.JSON(utils.SuccessJson(data))
}
