package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:@/go_rest_api_crud?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("Connection Refused ", err)
	} else {
		log.Println("Connection Successed")
	}
}
