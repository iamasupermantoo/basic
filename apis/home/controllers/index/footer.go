package index

import (
	"basic/apis/home/service/index"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

// FooterInfo
// @Summary	查询底部列表
// @Tags Basic
// @Success	200			{object}	dtoindex.FooterData
// @Router		/api/v1/footer [post]
func FooterInfo(c *fiber.Ctx) error {
	lang := c.Get("Accept-Language")
	originParseUrl, _ := url.Parse(c.Get("Origin"))

	data, err := index.FooterIndex(originParseUrl.Host, lang)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
