package main

import (
	"github.com/spf13/viper"
	"jwtcli/cli"
)

func main() {

	viper.SetEnvPrefix("jwt")
	viper.BindEnv("password")

	cli.RootCmd().Execute()
}
