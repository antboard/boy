/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"strings"

	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "当前目录开一个文件上传下载服务",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()
		r.POST("/upload", func(c *gin.Context) {
			file, err := c.FormFile("file")
			if err != nil {
				log.Println(err)
				c.String(500, "上传失败")
				return
			}
			c.String(200, "%s uploaded!", file.Filename)
			// 保存到本地
			err = c.SaveUploadedFile(file, file.Filename)
			if err != nil {
				log.Println(err)
				return
			}
		})
		r.Static("/download", "./")
		// 获取本机ip
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("使用以下地址访问服务")
		for _, addr := range addrs {
			if strings.Contains(addr.String(), ".") {
				if strings.Contains(addr.String(), "127") {
					continue
				}
				ip := strings.Split(addr.String(), "/")[0]
				fmt.Println("下载: http://" + ip + ":6666/download/当前路径文件名")
				fmt.Println("上传: " + `curl -X POST -F 'file=@文件名'  http://` + ip + ":6666/upload")
			}
		}
		err = r.Run(":6666")
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
}
