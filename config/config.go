package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Database 数据库配置
type Database struct {
	Network string `yaml:"Network"` //	网络
	Server  string `yaml:"Server"`  //	服务地址
	Port    int    `yaml:"Port"`    //	端口
	User    string `yaml:"User"`    //	用户名
	Pass    string `yaml:"Pass"`    //	密码
	DbName  string `yaml:"DbName"`  //	数据库
}

// Redis Redis配置
type Redis struct {
	Network         string `yaml:"Network"`         //	网络
	Server          string `yaml:"Server"`          //	服务器地址
	Port            int    `yaml:"Port"`            //	端口
	Pass            string `yaml:"Pass"`            //	密码
	DbName          int    `yaml:"DbName"`          //	数据库
	ConnectTimeout  int64  `yaml:"ConnectTimeout"`  //	连接超时时间
	ReadTimeout     int64  `yaml:"ReadTimeout"`     //	读取超时时间
	WriteTimeout    int64  `yaml:"WriteTimeout"`    //	写入超时时间
	MaxOpenConn     int    `yaml:"MaxOpenConn"`     // 	设置最大连接数
	ConnMaxIdleTime int64  `yaml:"ConnMaxIdleTime"` // 	空闲连接超时
	MaxIdleConn     int    `yaml:"MaxIdleConn"`     // 	最大空闲连接数
	Wait            bool   `yaml:"Wait"`            // 	如果超过最大连接数是否等待
}

type Basic struct {
	Name         string `yaml:"Name"`         //	项目名称
	Debug        bool   `yaml:"Debug"`        //	Debug
	Port         string `yaml:"Port"`         //	启动端口
	AdmDomain    string `yaml:"AdmDomain"`    //	后台接口域名
	ServerHeader string `yaml:"ServerHeader"` //	Header 名称
	Prefork      bool   `yaml:"Prefork"`      //	是否开启子进程
}

// Config 配置文件
type Config struct {
	Basic    *Basic    `yaml:"Basic"`    //	APP配置
	Database *Database `yaml:"Database"` //	数据库配置
	Redis    *Redis    `yaml:"Redis"`    //	缓存配置
}

// ReadYamlConfigFile 读取Yaml配置文件
func ReadYamlConfigFile() *Config {
	var conf *Config
	conFile, err := os.ReadFile("./config/app.yaml")
	if err != nil {
		panic(err)
	}

	//	映射配置文件
	err = yaml.Unmarshal(conFile, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
