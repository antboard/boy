package project

import (
	"boy/pkg/files"
	_ "embed"
	"log"
	"os"
)

// 创建配置的实现
// 1. 一个简单的map 实现kv读写.
// 2. 一个etcd的实现,实现kv读写

//go:embed files/config/simple_config.go
var mapConfig string

var configFiles map[string]string

func init() {
	configFiles = make(map[string]string)
	configFiles["pkg/config/simple_config.go"] = mapConfig
}

func CreateSimpleConfigFile() error {
	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}
	err2 := files.CreateDirAndFile(curDir, configFiles)
	if err2 != nil {
		return err2
	}
	return nil
}
