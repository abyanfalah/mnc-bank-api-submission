package main

import (
	"fmt"
	"mnc-bank-api/config"
	"mnc-bank-api/server"
	"mnc-bank-api/utils/migration"
	"os"
	"strings"
)

const (
	API_HOST = "localhost"
	API_PORT = "8000"

	APP_NAME = "MNC_BANK_API_TEST"
)

func main() {
	migration.Migrate()
	setEnv()
	viewConfigs()

	server.NewAppServer().Run()

}

func setEnv() {
	os.Setenv("API_HOST", API_HOST)
	os.Setenv("API_PORT", API_PORT)

	os.Setenv("APP_NAME", APP_NAME)

}

func viewConfigs() {
	fmt.Println(strings.Repeat("=", 50))

	config := config.NewConfig()
	fmt.Println("api config (port maybe auto set):")
	fmt.Println("host:", config.ApiConfig.Host)
	fmt.Println("port:", config.ApiConfig.Port)

	fmt.Println(strings.Repeat("=", 50))
}
