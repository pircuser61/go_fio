package config

import (
	"fmt"
	"os"
)

func GetHTTPPort() string {
	return getEnv("GO_FIO_REST_PORT", ":8088")
}

func GetConnectionString() string {
	dbname := getEnv("POSTGRES_DB", "postgres")
	host := getEnv("POSTGRES_HOST", "localhost")
	port := getEnv("POSTGRES_PORT", "5433")
	user := getEnv("POSTGRES_USER", "user")
	passw := getEnv("POSTGRES_PASSWORD", "1234")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, passw, dbname)
}

func getEnv(name string, defaultValue string) string {
	val, ok := os.LookupEnv(name)
	if ok {
		return val
	}
	return defaultValue
}
