package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type Config struct {
	RouterEngine   *gin.Engine
	ApiUrl         string
	DataSourceName string
}

func NewConfig() *Config {
	config := new(Config)
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DBNAME")
	config.RouterEngine = gin.Default()
	config.ApiUrl = "localhost:3030"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	config.DataSourceName = dsn
	return config
}
