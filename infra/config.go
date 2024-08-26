package infra

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/kchxng/go-oracle-api/utils/logs"
)

func InitConfig() {
	err := godotenv.Load(".env")
	if err == nil {
		logs.Info("Loading from .env file...")
	}
}

func InitTimezone() {
	ict, err := time.LoadLocation(GetENV("app.timezone"))
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func DatabaseParams() (string, int, string, string, string) {
	host := GetENV("database_host")
	port := GetENV("database_port")
	serviceName := GetENV("database_service_name")
	username := GetENV("database_username")
	password := GetENV("database_password")
	portInt, err := strconv.Atoi(port)

	if err != nil {
		panic(err)
	}
	return host, portInt, serviceName, username, password

}

func DatabaseServiceName() string {
	databaseName := GetENV("service_name")
	return databaseName

}

func GetENV(k string) string {
	return os.Getenv(k)
}
