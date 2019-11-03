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
	db = utils.ConnectDB("root", "", "localhost", "test1", 3306)
	db.AutoMigrate(&Account{}) // risk?
}

func generateUsername() string {
	return utils.GenerateRandomString(10)
}

func encryptText(text string) string {
	return utils.GenerateMD5String(text)
}

func ConnectDB(uname, pwd, host, dbName string, port int) {
	db = utils.ConnectDB(uname, pwd, host, dbName, port)
}

func MigrateAccount() {
	db.AutoMigrate(&Account{})
}

func DropTable(name string) error {
	return db.DropTable(name).Error
}

func HasTable(name string) bool {
	return db.HasTable(name)
}

func CreateTableAccount() error {
	if err := db.CreateTable(&Account{}).Error; err != nil {
		return err
	}
	return nil
}

func DropTableAccount() error {
	if err := db.DropTable(&Account{}).Error; err != nil {
		return err
	}
	return nil
}

func HasTableAccount() bool {
	return db.HasTable(&Account{})
}

