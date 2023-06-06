package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Config struct {
	RouterEngine   *gin.Engine
	ApiUrl         string
	DataSourceName string
}

func NewConfig() *Config {
	config := new(Config)
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DBNAME")
	config.RouterEngine = gin.Default()
	config.ApiUrl = os.Getenv("API_URL")
	if config.ApiUrl == "" {
		config.ApiUrl = ":3030"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	log.Println(dsn)
	config.DataSourceName = dsn
	return config
}
