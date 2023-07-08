/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
		err := makeCall(args) // 先构建一下
		if err != nil {
			return
		}
		runCall(args)
	},
}

func runCall(args []string) {
	var execCmd *exec.Cmd
	execCmd = exec.Command("make", "run")
	for i := 0; i < len(args); i++ {
		if os.DirFS(args[i]) != nil {
			execCmd.Dir = "./" + args[i]
		}
	}

	execCmd.Stderr = os.Stderr
	execCmd.Stdout = os.Stdout
	err := execCmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func init() {
	rootCmd.AddCommand(runCmd)
}
