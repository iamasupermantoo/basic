package level

import (
	"basic/apis/common/validator"
	dtousers "basic/apis/home/dto/users"
	"basic/apis/home/service/users"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Create
// @Summary	购买等级
// @Tags		User
// @Param		id	query		int		true	"等级ID"
// @Success	200	{string}	string	"ok"
// @Router		/api/v1/user/level/order [post]
func Create(c *fiber.Ctx) error {
	params := &dtousers.LevelOrderParams{}
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

	//	调用服务层
	data, err := users.LevelOrderCreate(adminId, userId, c.Get("Accept-Language"), params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	返回数据
	return c.JSON(utils.SuccessJson(data))
}
