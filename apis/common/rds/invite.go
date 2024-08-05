package rds

import (
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"encoding/json"

	"github.com/gomodule/redigo/redis"
)

const (
	// RedisInviteCode 缓存邀请码
	RedisInviteCode = "RedisInviteCode"
)

// RedisFindInviteCodeInfo 获取邀请码信息
func RedisFindInviteCodeInfo(rds redis.Conn, code string) (*models.Invite, error) {
	//	获取缓存是否存在邀请信息
	inviteInfo, err := GetRedisInviteCodeInfo(rds, code)
	if err != nil {
		inviteInfo = &models.Invite{}
		result := database.Db.Where("code = ?", code).Find(&inviteInfo)
		if result.Error != nil {
			return nil, result.Error
		}

		SetRedisInviteCodeInfo(rds, code, inviteInfo)
	}

	return inviteInfo, nil
}

// SetRedisInviteCodeInfo 设置code信息缓存
func SetRedisInviteCodeInfo(rds redis.Conn, code string, inviteInfo *models.Invite) {
	inviteInfoBytes, _ := json.Marshal(inviteInfo)
	_, _ = rds.Do("HSET", RedisInviteCode, code, inviteInfoBytes)
}

// DelRedisInviteCodeInfo 删除邀请码信息
func DelRedisInviteCodeInfo(rds redis.Conn, code string) {
	_, _ = rds.Do("HDEL", RedisInviteCode, code)
}

// DelRedisInvite 删除所有邀请缓存
func DelRedisInvite(rds redis.Conn) {
	_, _ = rds.Do("DEL", RedisInviteCode)
}

// GetRedisInviteCodeInfo 获取code信息缓存
func GetRedisInviteCodeInfo(rds redis.Conn, code string) (*models.Invite, error) {
	inviteInfoBytes, err := redis.Bytes(rds.Do("HGET", RedisInviteCode, code))
	if err != nil {
		return nil, err
	}

	inviteInfo := &models.Invite{}
	_ = json.Unmarshal(inviteInfoBytes, &inviteInfo)
	return inviteInfo, nil
}

// DeleteRedisInvite 删除邀请缓存
func DeleteRedisInvite(ids []int) error {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	for _, v := range ids {
		inviteInfo := &models.Invite{}
		result := database.Db.Model(inviteInfo).Where("id = ?", v).Take(inviteInfo)
		if result.Error == nil {
			DelRedisInviteCodeInfo(rdsConn, inviteInfo.Code)
		}
	}
	return nil
}
