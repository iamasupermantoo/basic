package index

import (
	"basic/apis/home/service/index"
	"basic/utils"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

// HelpersInfo
// @Summary	查询帮助列表
// @Tags Basic
// @Success	200			{object}	dtoindex.HomeHelpersInfo
// @Router		/api/v1/helpers [post]
func HelpersInfo(c *fiber.Ctx) error {
	originParseUrl, _ := url.Parse(c.Get("Origin"))
	lang := c.Get("Accept-Language")

	data, err := index.HelpersInfo(originParseUrl.Host, lang)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
