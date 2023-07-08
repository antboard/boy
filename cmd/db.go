/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"boy/internal/project"
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "添加数据库",
	Long:  `参数 mysql postgresql `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		switch args[0] {
		case "mysql":
			project.CreateMysqlFile()
		case "postgresql":
			project.CreatePostgresqlFile()
		}
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
}
