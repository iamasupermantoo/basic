package index

import (
	"basic/apis/common/service/index"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Upload 上传文件
func Upload(c *fiber.Ctx) error {
	//解析文件数据
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(utils.NewResponseJson().Error(err.Error()))
	}

	//获取所有文件
	files := form.File["file"]
	var filepaths []string

	//上传文件
	for _, file := range files {
		saveFilePath, err := index.Upload(file)
		if err != nil {
			return err
		}
		filepaths = append(filepaths, saveFilePath.(string))
	}

	return c.JSON(utils.NewResponseJson().Success(filepaths))
}
