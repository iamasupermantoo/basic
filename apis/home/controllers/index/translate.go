package index

import (
	"basic/apis/common/validator"
	dtohomeindex "basic/apis/home/dto/index"
	"basic/apis/home/service/index"
	"basic/utils"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

// TranslateList
// @Summary	查询前台翻译
// @Tags Basic
// @Param lang body string true "语言标识"
// @Success	200			{object}	dto.TranslateOptions
// @Router		/api/v1/translate [post]
func TranslateList(c *fiber.Ctx) error {
	params := &dtohomeindex.HomeTranslateParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	originParseUrl, _ := url.Parse(c.Get("Origin"))
	data, err := index.TranslateInfo(originParseUrl.Host, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
