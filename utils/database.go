package utils

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func ConnectDB(uname, pwd, host, dbname string, port int) *gorm.DB {
	args := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", uname,pwd, host, port, dbname)
	db, err := gorm.Open("mysql", args)
	if err != nil {
		// panic("failed to connect database" + err)
		panic(err)
	}
	return db
}
