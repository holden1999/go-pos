package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Config struct {
	RouterEngine   *gin.Engine
	ApiUrl         string
	DataSourceName string
}

func NewConfig() *Config {
	config := new(Config)
	dbHost := "127.0.0.1" //os.Getenv("DB_HOST")
	dbPort := "3306"      //os.Getenv("DB_PORT")
	dbUser := "root"      //os.Getenv("DB_USER")
	dbPassword := "root"  //os.Getenv("DB_PASSWORD")
	dbName := "pos"       //os.Getenv("DB_DBNAME")
	config.RouterEngine = gin.Default()
	config.ApiUrl = "localhost:3030"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	config.DataSourceName = dsn
	return config
}
