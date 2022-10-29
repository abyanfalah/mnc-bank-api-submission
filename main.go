package main

import (
	"fmt"
	"mnc-bank-api/config"
	"mnc-bank-api/server"
	"os"
	"strings"
)

const (
	DB_HOST   = "localhost"
	DB_PORT   = "5432"
	DB_USER   = "postgres"
	DB_PASS   = "12345"
	DB_NAME   = "mnc_bank_test"
	DB_DRIVER = "postgres"

	API_HOST = "localhost"
	API_PORT = "8000"

	APP_NAME = "MNC_BANK_API_TEST"
)

func main() {
	setEnv()
	viewConfigs()

	server.NewAppServer().Run()

	// try.Run()
}

func setEnv() {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Auto setting environment variables")
	fmt.Println("You can disable this feature on main.go")

	// os.Setenv("DB_HOST", DB_HOST)
	// os.Setenv("DB_PORT", DB_PORT)
	// os.Setenv("DB_USER", DB_USER)
	// os.Setenv("DB_PASS", DB_PASS)
	// os.Setenv("DB_NAME", DB_NAME)
	// os.Setenv("DB_DRIVER", DB_DRIVER)

	os.Setenv("API_HOST", API_HOST)
	os.Setenv("API_PORT", API_PORT)

	os.Setenv("APP_NAME", APP_NAME)

	fmt.Println("Setting finished")
	fmt.Println(strings.Repeat("=", 50))

}

func viewConfigs() {

	config := config.NewConfig()

	fmt.Println("configs: ")
	fmt.Println("api config (port maybe auto set):", config.ApiConfig)

	fmt.Println(strings.Repeat("=", 50))

	fmt.Println()

}
