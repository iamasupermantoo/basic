package admins

import (
	"basic/models"
	"basic/module/database"
	"basic/utils/views"
)

// GetRouterNameList 获取所有路由名称
func GetRouterNameList() []*views.InputOptions {
	var authChildList []*models.AuthChild
	if result := database.Db.Model(&models.AuthChild{}).
		Where("type = ?", models.AuthChildTypeRouteNameRoute).
		Find(&authChildList); result.Error != nil {
		return nil
	}

	data := make([]*views.InputOptions, 0)
	for i := 0; i < len(authChildList); i++ {
		data = append(data, &views.InputOptions{
			Label: authChildList[i].Parent,
			Value: authChildList[i].Parent,
		})
	}
	return data
}
