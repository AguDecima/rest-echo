package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Register some standard stuff
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Register some standard stuff
)

// DB is...
var DB *gorm.DB

// InitDB is...
func InitDB() {
	var err error
	dataSourceName := "root:root1234@tcp(127.0.0.1:3306)/tests?charset=utf8&parseTime=True"
	DB, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	DB.Debug().DropTableIfExists(&User{})
	//Drops table if already exists
	DB.Debug().AutoMigrate(&User{})
}
