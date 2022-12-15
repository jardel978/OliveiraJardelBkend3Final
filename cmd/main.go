package main

import (
	"OliveiraJardelBkend3Final/cmd/server"
	"OliveiraJardelBkend3Final/internal/configs"
)

// @title CheckPoint - Backend 3 Final
// @version 1.0
// @description This API Handle Products.
// @termsOfService

// @contact.name Jardel e Nelson
// @contact.url https://www.developersprofissa.com.ko/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	configs.Load()
	configs.StartGormDb()
	server := server.NewServer()
	server.InitSwag()
	server.Run()
}

//package main
//
//import (
//"fmt"
//"gorm.io/driver/postgres"
//"gorm.io/gorm"
//"time"
//)
//
//type Animal struct {
//	AnimalID int64     `gorm:"column:beast_id"`
//	Birthday time.Time `gorm:"column:day_of_the_beast"`
//	Age      int64     `gorm:"column:age_of_the_beast"`
//}
//
//func main() {
//	dsn := "host=postgres user=postgres password=***** dbname=postgres port=5432"
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	db.AutoMigrate(&Animal{})
//
//	result, _ := db.Debug().Migrator().ColumnTypes(&Animal{})
//	for _, v := range result {
//		fmt.Println(v.Name())
//	}
//}
