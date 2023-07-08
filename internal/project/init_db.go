package project

import (
	"boy/pkg/files"
	_ "embed"
	"log"
	"os"
)

// 添加数据库

//go:embed files/db/mysql.go
var mysql string

//go:embed files/db/postgresql.go
var postgresql string

var dbMysqlFiles map[string]string
var dbPostgresqlFiles map[string]string

func init() {
	dbMysqlFiles = make(map[string]string)
	dbMysqlFiles["pkg/db/mysql.go"] = mysql

	dbPostgresqlFiles = make(map[string]string)
	dbPostgresqlFiles["pkg/db/postgresql.go"] = postgresql
}

func CreateMysqlFile() error {
	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}
	err2 := files.CreateDirAndFile(curDir, dbMysqlFiles)
	if err2 != nil {
		return err2
	}
	return nil
}

func CreatePostgresqlFile() error {
	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}
	err2 := files.CreateDirAndFile(curDir, dbPostgresqlFiles)
	if err2 != nil {
		return err2
	}
	return nil
}
