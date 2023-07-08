package files

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func CreateDirAndFile(curDir string, fileList map[string]string) (err error) {
	for path, data := range fileList {
		// 创建目录
		if strings.HasSuffix(path, "/") {
			// 纯目录
			err := exec.Command("mkdir", "-p", curDir+"/"+path).Run()
			if err != nil {
				log.Println(err)
				return err
			}
		} else {
			// 创建目录
			if strings.Index(path, "/") >= 0 {
				// 目录文件混合, 先创建目录, 再创建文件
				idx := strings.LastIndex(path, "/")
				//fileName := path[idx+1:]
				//log.Println(fileName)
				// 创建目录
				err := exec.Command("mkdir", "-p", curDir+"/"+path[:idx]).Run()
				if err != nil {
					log.Println(err)
					return err
				}
			}

			// 创建文件
			_, err = os.Stat(curDir + "/" + path)
			if err != nil {
				err := os.WriteFile(curDir+"/"+path, []byte(data), 0666)
				if err != nil {
					log.Println(err)
					return err
				}
			}
		}
	}
	return nil
}
