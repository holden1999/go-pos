package manager

import (
	"database/sql"
	"go-pos/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type InfraManager interface {
	SqlDb() *gorm.DB
}

type infraManager struct {
	db *gorm.DB
}

func NewInfra(config *config.Config) InfraManager {
	resource := initDb(config.DataSourceName)
	return &infraManager{db: resource}
}

func (i infraManager) SqlDb() *gorm.DB {
	return i.db
}

func initDb(dataSourceName string) *gorm.DB {
	sqlDB, _ := sql.Open("mysql", dataSourceName)
	conn, _ := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}))
	return conn
}
