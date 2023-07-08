package project

import (
	"boy/pkg/files"
	_ "embed"
	"log"
	"os"
)

//go:embed files/main.go
var MainGo string

//go:embed files/Makefile
var makefile string

var baseFiles map[string]string

func init() {
	baseFiles = make(map[string]string)
	baseFiles["cmd/service/main.go"] = MainGo
	baseFiles["internal/feature1/dto.go"] = "package feature1\n\n// 文件结构杜绝了其他包引用dto, model是基础数据结构."
	baseFiles["internal/feature1/handle.go"] = "package feature1\n\n// 仅处理一种功能. 如果引用了其他功能, 那就独立搞一个聚合类的功能."
	baseFiles["internal/model/model.go"] = "package model\n\n// 数据库表的一个结构体, 与, 增删改查等.\n// 非数据库,直接使用pb的模型."
	baseFiles["internal/server.go"] = "package internal\n\n// grpc服务实现 /gin 路由的代码\n"
	baseFiles["output/"] = ""
	baseFiles["pkg/"] = ""
	baseFiles["Makefile"] = makefile
	baseFiles[".gitignore"] = "output\n"
}

func CreateBaseFile() error {
	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}
	err2 := files.CreateDirAndFile(curDir, baseFiles)
	if err2 != nil {
		return err2
	}
	return nil
}
