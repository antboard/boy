/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "编译当前目录的工程",
	Long:  `linux: 编译称linux elf格式, 子目录名称: 进入子目录执行make`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("make called")
		_ = makeCall(args)

	},
}

func makeCall(args []string) error {
	var execCmd *exec.Cmd
	execCmd = exec.Command("make")
	var env []string
	for i := 0; i < len(args); i++ {
		if args[i] == "linux" {
			env = append(env, "GOOS=linux")
			continue
		}
		if args[i] == "amd64" {
			env = append(env, "GOARCH=amd64")
			continue
		}

		if os.DirFS(args[i]) != nil {
			execCmd.Dir = "./" + args[i]
		}
	}
	if len(env) > 0 {
		execCmd.Env = append(os.Environ(), env...)
	}
	execCmd.Stderr = os.Stderr
	execCmd.Stdout = os.Stdout
	err := execCmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(makeCmd)
}
