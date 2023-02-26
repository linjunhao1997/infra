/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var goInputFilePath string
var goOutputFilePath string

// gozeroapiCmd represents the gozeroapi command
var gozeroapiCmd = &cobra.Command{
	Use:   "gozeroapi",
	Short: "生成gozero的api文件",
	Long:  `通过解析go文件的结构体的tag生成对应的api文件,可选择将api内容输出到文件或控制台中`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gozeroapi called")
	},
}

func init() {
	gozeroapiCmd.Flags().StringVarP(&goInputFilePath, "input", "in", "./", "go文件路径")
	gozeroapiCmd.Flags().StringVarP(&goInputFilePath, "input", "in", "./", "go文件路径")

	rootCmd.AddCommand(gozeroapiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gozeroapiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gozeroapiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
