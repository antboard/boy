package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

type PostgresqlKv interface {
	Get(key string) string // 获取配置
}

var postgresqls map[string]*gorm.DB

func InitPostgresql(names []string, kv PostgresqlKv) {
	postgresqls = make(map[string]*gorm.DB)
	for _, name := range names {
		dbName := kv.Get(name + "_db_name")
		dbUser := kv.Get(name + "_db_user")
		dbPassword := kv.Get(name + "_db_password")
		dbAddress := kv.Get(name + "_db_address")

		list := strings.Split(dbAddress, ":")
		host := list[0]
		port := "5432"
		if len(list) > 1 {
			port = list[1]
		}

		source := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			host, dbUser, dbPassword,
			dbName, port)
		db, err := gorm.Open(postgres.New(postgres.Config{DSN: source, PreferSimpleProtocol: true}), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		postgresqls[name] = db
	}
}

func GetPostgresql(name string) *gorm.DB {
	return postgresqls[name]
}
