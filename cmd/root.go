/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "boy",
	Short: "一个 golang语言的工程管理工具",
	Long:  `可以很容易的创建一个工程, 集成一些常用的包, 欢迎提出建设意见`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	fmt.Println("    __\n   / /_  ____  __  __\n  / __ \\/ __ \\/ / / /\n / /_/ / /_/ / /_/ /\n/_.___/\\____/\\__, /\n            /____/")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
