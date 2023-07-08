/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"boy/internal/project"
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化工程文件",
	Long:  `默认初始化工程文件, 带一个实例. `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("在当前目录初始化一个默认的工程结构")
		project.CreateBaseFile()
		fmt.Println("Please run go mod init <MODNAME> and go mod tidy")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
