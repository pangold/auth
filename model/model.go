package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"../utils"
)

var (
	db *gorm.DB
)

func init() {
	db = utils.ConnectDB("root", "88888888", "localhost", "test1", 3306)
	db.AutoMigrate(&Account{}) // risk?
}

func generateUsername() string {
	return utils.GenerateRandomString(10)
}

func encryptText(text string) string {
	return utils.GenerateMD5String(text)
}
