package common

import "github.com/asim/go-micro/v3/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

//获取mysql配置

func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlconfig := &MysqlConfig{}
	config.Get(path...).Scan(mysqlconfig)
	return mysqlconfig
}
