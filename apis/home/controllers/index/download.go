package index

import (
	"basic/apis/home/service/index"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

// DownloadInfo
// @Summary	查询下载信息
// @Tags Basic
// @Success	200			{object}	dtoindex.HomeDownLoadInfo
// @Router		/api/v1/download [post]
func DownloadInfo(c *fiber.Ctx) error {
	originParseUrl, _ := url.Parse(c.Get("Origin"))

	data, err := index.DownloadInfo(originParseUrl.Host)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
