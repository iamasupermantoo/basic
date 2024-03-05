package index

import (
	"basic/apis/common/rds"
	dtoindex "basic/apis/home/dto/index"
	"basic/module/cache"
	"github.com/goccy/go-json"
)

// DownloadInfo 获取下载列表
func DownloadInfo(host string) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	域名获取管理ID
	adminId := rds.RedisFindDomainToAdminId(rdsConn, host)
	value := rds.RedisFindAdminSettingField(rdsConn, adminId, "download")

	data := &dtoindex.HomeDownLoadInfo{}
	if err := json.Unmarshal([]byte(value), data); err != nil {
		return nil, err
	}
	return data, nil
}
