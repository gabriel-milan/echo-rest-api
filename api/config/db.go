package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var DefaultDBUser string = os.Getenv("DB_USER")
var DefaultDBPassword string = os.Getenv("DB_PASSWORD")
var DefaultDBName string = os.Getenv("DB_NAME")
var DefaultDBHost string = os.Getenv("DB_HOST")
var DefaultDBPort string = os.Getenv("DB_PORT")
var DefaultSSLMode string = os.Getenv("DB_SSL_ENABLED")

func GetDefaultDBUser() string {
	if DefaultDBUser == "" {
		return "postgres"
	}
	return DefaultDBUser
}

func GetDefaultDBPassword() string {
	if DefaultDBPassword == "" {
		return "R9EXkbAnCPkF4tpm"
	}
	return DefaultDBPassword
}

func GetDefaultDBName() string {
	if DefaultDBName == "" {
		return "postgres"
	}
	return DefaultDBName
}

func GetDefaultDBHost() string {
	if DefaultDBHost == "" {
		return "postgres"
	}
	return DefaultDBHost
}

func GetDefaultDBPort() string {
	if DefaultDBPort == "" {
		return "5432"
	}
	return DefaultDBPort
}

func GetDefaultSSLMode() string {
	if DefaultSSLMode == "" {
		return "disable"
	}
	enabled, err := strconv.ParseBool(DefaultSSLMode)
	if err != nil {
		log.Fatal(err)
	}
	if enabled {
		return "enable"
	}
	return "disable"
}

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		GetDefaultDBHost(),
		GetDefaultDBPort(),
		GetDefaultDBUser(),
		GetDefaultDBName(),
		GetDefaultDBPassword(),
		GetDefaultSSLMode(),
	)
	return dataBase
}
