package bill

import (
	"basic/apis/common/validator"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/apis/home/service/wallets"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index
// @Summary	用户账单列表
// @Tags Wallet
// @Param		types		body	[]int false "类型 1充值类型 11提现类型 21购买 51收益 61奖励"
// @Param		createdAt		body	utils.RangeDatePicker false "时间范围"
// @Param		pagination		body	utils.Pagination   false "分页"
// @Success		200			{object}	utils.ResponseJson{data=[]dtowallets.HomeBillIndexData}
// @Router		/api/v1/wallets/bill/index [post]
func Index(c *fiber.Ctx) error {
	params := &dtowallets.BillIndexParams{}
	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	err := validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := wallets.BillIndex(adminId, userId, c.Get("Accept-Language"), params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
