package index

import (
	"basic/apis/common/rds"
	dtoindex "basic/apis/home/dto/index"
	"basic/module/cache"
)

// TranslateInfo 获取字典列表
func TranslateInfo(host string, params *dtoindex.HomeTranslateParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	域名获取管理ID
	adminId := rds.RedisFindDomainToAdminId(rdsConn, host)
	data := rds.RedisFindHomeTranslateLang(rdsConn, adminId, params.Lang)
	return data, nil
}
