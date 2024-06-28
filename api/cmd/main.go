/*
Copyright © 2024 NAME HERE boscardinvinicius@gmail.com
*/
package main

import (
	cmd "github.com/booscaaa/go-gemini-gdg/api/cmd/cmds"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(".env not exists")
	}
}

func main() {
	cmd.Execute()
}
