package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var GormOpen *gorm.DB

func init() {
	var err error
	GormOpen, err = gorm.Open("mysql", "root:12345600@tcp(127.0.0.1:3307)/BUYDB?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if GormOpen.Error != nil {
		fmt.Printf("database error %v", GormOpen.Error)
	}
}
