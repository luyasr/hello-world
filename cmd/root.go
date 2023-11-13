/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/luyasr/hello-world/pkg/ioc"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hello world",
	Short: "一个简单的hello world程序",
	Long:  "使用grpc和gin框架实现的hello world demo",

	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.OnInitialize(InitIoc)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func InitIoc() {
	// 初始化controller
	if err := ioc.Controller().Init(); err != nil {
		panic(err)
	}

	// 初始化api server
	if err := ioc.HttpHandler().Init(); err != nil {
		panic(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.user.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
