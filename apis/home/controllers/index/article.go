package index

import (
	"basic/apis/common/validator"
	dtoindex "basic/apis/home/dto/index"
	"basic/apis/home/service/index"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

// ArticleIndex 获取文章列表
// @Summary 获取文章列表
// @Tags Basic
// @Param types body []int false "类型数组： 1帮助中心类型 2底部类型 10基础文章"
// @Success 200 {object} dtoindex.HomeArticleInfo
// @Router /api/v1/article/index [post]
func ArticleIndex(c *fiber.Ctx) error {
	params := &dtoindex.ArticleIndexParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	originParseUrl, _ := url.Parse(c.Get("Origin"))
	lang := c.Get("Accept-Language")
	data, err := index.ArticleIndex(originParseUrl.Host, lang, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}

// ArticleInfo 获取文章信息
// @Summary 获取文章信息
// @Tags Basic
// @Param id body int true "文章Id"
// @Success 200 {object} dtoindex.HomeArticleInfo
// @Router /api/v1/article/info [post]
func ArticleInfo(c *fiber.Ctx) error {
	params := &dtoindex.ArticleInfoParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	lang := c.Get("Accept-Language")
	data, err := index.ArticleInfo(lang, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
