package main

func main() {
	NewApiServer().Run()
}

//func main() {
//	dsn := "root:root@tcp(127.0.0.1:3306)/pos?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
//	// Migrate the schema
//	err = db.AutoMigrate(&model.Cashier{}, &model.Category{}, &model.Discount{}, &model.Order{}, &model.Payment{}, &model.Product{})
//	if err != nil {
//		return
//	}
//}
