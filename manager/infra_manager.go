package manager

import (
	"database/sql"
	"go-pos/config"
	"go-pos/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type InfraManager interface {
	SqlDb() *gorm.DB
}

type infraManager struct {
	db *gorm.DB
}

func (i infraManager) SqlDb() *gorm.DB {
	err := i.db.AutoMigrate(&model.Cashier{}, &model.Category{}, &model.Credential{}, &model.Discount{}, &model.Order{}, &model.Payment{}, model.Product{})
	if err != nil {
		return nil
	}
	return i.db
}

func initDb(dataSourceName string) *gorm.DB {
	sqlDB, _ := sql.Open("mysql", dataSourceName)
	conn, _ := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})

	return conn
}

func NewInfra(config *config.Config) InfraManager {
	resource := initDb(config.DataSourceName)
	return &infraManager{
		db: resource,
	}
}
