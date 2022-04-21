package config

import "github.com/gin-gonic/gin"

type Config struct {
	RouterEngine   *gin.Engine
	ApiUrl         string
	DataSourceName string
}

func NewConfig() *Config {
	config := new(Config)
	config.RouterEngine = gin.Default()
	config.ApiUrl = "localhost:3030"
	config.DataSourceName = "root:root@tcp(127.0.0.1:3306)/pos?charset=utf8mb4&parseTime=True&loc=Local"
	return config
}
