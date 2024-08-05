package models

import (
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"
	"crypto/rsa"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gomodule/redigo/redis"
)

const (
	// RedisTokenRsa Redis Rsa私钥
	RedisTokenRsa = "TokenRsa"

	// TokenExpired Token过期时间
	TokenExpired = 86400 * 30

	// ServiceAdminName 后台项目名称
	ServiceAdminName = "admin"

	//	ServiceHomeName 前台项目名称
	ServiceHomeName = "home"

	// RedisTokenName Token名称
	RedisTokenName = "_Token"
)

// TokenParams Token参数
type TokenParams struct {
	Name    string `json:"name"`    //	设备名称
	Token   string `json:"token"`   //	Token
	Expired int64  `json:"expired"` //	过期时间
}

// RsaSetting Rsa配置
type RsaSetting struct {
	PriKey string `json:"priKey"` //	私钥
	PubKey string `json:"pubKey"` //	公钥
}

// AdminUserSuperData 超级管理数据
type AdminUserSuperData struct {
	RsaList map[string]*RsaSetting `json:"rsaList"` //	Rsa列表
}

// NewServiceAdminToken 生成后端服务Token
func NewServiceAdminToken(adminId int) string {
	return NewServiceToken(ServiceAdminName, adminId, 0, TokenExpired)
}

// NewServiceHomeToken 生成前端服务Token
func NewServiceHomeToken(adminId, userId int) string {
	return NewServiceToken(ServiceHomeName, adminId, userId, TokenExpired)
}

// NewServiceToken 生成服务Token
func NewServiceToken(serviceName string, adminId, userId, expired int) string {
	privateKey, _ := GetServiceRsaPrivate(serviceName)
	claims := jwt.MapClaims{
		"adminId": adminId,
		"userId":  userId,
		"exp":     time.Now().Add(time.Second * time.Duration(expired)).Unix(),
	}

	//	生成Token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	s, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	//	设置缓存Token
	SetRedisTokenParams(adminId, userId, s, int64(expired))
	return s
}

// GetSuperAdminUserData 获取超级管理数据
func GetSuperAdminUserData() *AdminUserSuperData {
	adminUserInfo := &AdminUser{}
	database.Db.First(&adminUserInfo, SuperAdminId)

	adminUSerData := &AdminUserSuperData{}
	_ = json.Unmarshal([]byte(adminUserInfo.Data), &adminUSerData)
	return adminUSerData
}

// GetServiceRsaPrivate 获取服务私钥
func GetServiceRsaPrivate(serviceName string) (*rsa.PrivateKey, error) {
	rds := cache.Rds.Get()
	defer rds.Close()

	privateKeyStr := RedisGetServiceRsaPrivate(rds, serviceName)
	if privateKeyStr == "" {
		//	数据库中查询
		adminSuperData := GetSuperAdminUserData()
		privateKeyStr = adminSuperData.RsaList[serviceName].PriKey
		RedsiSetServiceRsaPrivate(rds, serviceName, privateKeyStr)
	}

	//	解析私钥
	privateKey, err := utils.ParseRsaPrivateStringData(privateKeyStr)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// RedisGetServiceRsaPrivate 获取缓存服务Rsa 私钥
func RedisGetServiceRsaPrivate(rds redis.Conn, serviceName string) string {
	privateKeyStr, _ := redis.String(rds.Do("HGET", RedisTokenRsa, serviceName))
	return privateKeyStr
}

// RedsiSetServiceRsaPrivate设置服务Rsa 私钥
func RedsiSetServiceRsaPrivate(rds redis.Conn, serviceName string, privateKeyStr string) {
	_, _ = rds.Do("HSET", RedisTokenRsa, serviceName, privateKeyStr)
}

// Token是否存在
func RedisTokenParams(c *fiber.Ctx) (int, int, bool) {
	rds := cache.Rds.Get()
	defer rds.Close()

	//	获取缓存中的Toke
	adminId, userId := utils.GetContextClaims(c)
	if adminId > 0 || userId > 0 {
		userTokenName := strconv.FormatInt(int64(adminId), 10) + "-" + strconv.FormatInt(int64(userId), 10)
		tokenParamsBytes, err := redis.Bytes(rds.Do("HGET", RedisTokenName, userTokenName))
		if err == nil {
			tokenParams := &TokenParams{}
			_ = json.Unmarshal(tokenParamsBytes, &tokenParams)
			if tokenParams.Token == utils.GetHeadersToken(c) {
				return adminId, userId, true
			}
		}
	}

	return 0, 0, false
}

// SetRedisTokenParams 缓存Token参数
func SetRedisTokenParams(adminId, userId int, token string, expired int64) {
	rds := cache.Rds.Get()
	defer rds.Close()

	//	唯一TokenKey
	userTokenName := strconv.FormatInt(int64(adminId), 10) + "-" + strconv.FormatInt(int64(userId), 10)
	tokenParams := &TokenParams{
		Name:    "",
		Token:   token,
		Expired: expired,
	}

	tokenParamsBytes, _ := json.Marshal(tokenParams)
	rds.Do("HSET", RedisTokenName, userTokenName, tokenParamsBytes)
}
