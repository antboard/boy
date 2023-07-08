package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlKv 通用配置接口, pkg中需要配置的包, 根据固定的格式拉取配置.
type MysqlKv interface {
	Get(key string) string // 获取配置
}

// 如果用户只有一个数据库,
var mysqls map[string]*gorm.DB

func InitMysql(names []string, kv MysqlKv) {
	mysqls = make(map[string]*gorm.DB)
	for _, name := range names {
		dbName := kv.Get(name + "_db_name")
		dbUser := kv.Get(name + "_db_user")
		dbPassword := kv.Get(name + "_db_password")
		dbAddress := kv.Get(name + "_db_address")
		db, err := gorm.Open(mysql.Open(dbUser+":"+dbPassword+"@tcp("+dbAddress+")/"+dbName+"?charset=utf8mb4&parseTime=True"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		mysqls[name] = db
	}
}

func GetMysql(name string) *gorm.DB {
	return mysqls[name]
}
