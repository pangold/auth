package model

import (
	"fmt"
	"errors"
	"regexp"
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	// You can use any one of UserId, Email or PhoneNumber to login
	// the UserId here is the same to the UserId of table User
	Username	string `json:"username"           gorm:"type:varchar(100);unique_index"`
	Email		string `json:"email"              gorm:"type:varchar(100)"`
	Phone		string `json:"phone"              gorm:"type:varchar(100)"`
	Password	string `json:"-"                  gorm:""`
	IsActivated	bool   `json:"is_activated"       gorm:"default:false"`	// Needs to activate when registerred
	IsLocked	bool   `json:"is_locked"          gorm:"default:false"`	// Input wrong userId & password 5 times
	IsEnabled	bool   `json:"is_enabled"         gorm:"default:false"`	// actively disabled, true after activated. (similar to deleted_at)
	Extended	string `json:"extended,omitempty" gorm:"type:TEXT"`
	// LastLogInAt	time.Time
	// LastLogOutAt	time.Time
	// LastLoggedIp	string
}

// Hooks
func (a *Account) BeforeSave() error {
	return a.IsValid()
}

func (a *Account) BeforeCreate() error {
	return a.IsValid()
}

func (a *Account) BeforeUpdate() error {
	return a.IsValid()
}

func (a *Account) BeforeDelete() error {
	return nil
}

// Column Verification
func (a *Account) IsUsernameValid() error {
	if len(a.Username) < 6 {
		return errors.New("username is too short.")
	}
	return nil
}

func (a *Account) IsEmailValid() error {
	// the first way to use regexp
	if a.Email == "" {
		// Email could be empty, so return directly
		return nil
	}
	if matched, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", a.Email); !matched {
		return errors.New("invalid email address")
	}
	return nil
}

func (a *Account) IsPhoneValid() error {
	// the second way to use regexp(similar)
	if a.Phone == "" {
		return nil
	}
	regular := "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$"
	reg := regexp.MustCompile(regular)
	if matched := reg.MatchString(a.Phone); !matched {
		return errors.New("invalid phone number")
	}
	return nil
}

func (a *Account) IsPasswordValid() error {
	// This is special, it needs to encrypt...
	// Now, this implementation is not correct, 
	// But it doesn't matter. It will always return nil
	if len(a.Password) < 6 {
		return errors.New("password is too short.")
	}
	return nil
}

func (a *Account) IsValid() error {
	if err := a.IsUsernameValid(); err != nil {
		return err
	}
	if err := a.IsEmailValid(); err != nil {
		return err
	}
	if err := a.IsPhoneValid(); err != nil {
		return err
	}
	if err := a.IsPasswordValid(); err != nil {
		return err
	}
	//? 
	//if len(a.Password) > 0 {
	//	a.Password = encryptText(a.Password)
	//}
	return nil
}

func (a *Account) IsStateValid() error {
	if !a.IsActivated {
		return errors.New("invalid account: in activated state")
	}
	if !a.IsEnabled {
		return errors.New("invalid account: in disabled state")
	}
	if a.IsLocked {
		return errors.New("invalid account: in locked state")
	}
	return nil
}

// JSON Translation
func (a Account) ToJsonString() string {
	res, err := json.Marshal(a)
	if err != nil {
		return fmt.Sprintf("{error: %s}", err.Error())
	}
	return string(res)
}

func (a *Account) FromJsonString(jsonString string) error {
	var accounts []Account
	err := json.Unmarshal([]byte(jsonString), &accounts)
	if err != nil {
		return err
	}
	*a = accounts[0]
	return nil
}

// Query
func GetAccounts() (accounts []Account, error) {
	if err := db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func GetValidAccounts() ([]Account, error) {
	return GetAccountsByState(Account{IsActivated: true, IsLocked: false, IsEnabled: true})
}

func GetAccountsByState(account Account) (accounts []Account, error) {
	// in those 3 states are bool: we can't get data if value = false
	if err := db.Where("is_activate = ? and is_locked = ? and is_enabled = ?", account.IsActivated, account.IsLocked, account.IsEnabled).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return
}

func GetActivatedAccounts(activated bool) (accounts []Account, error) {
	if err := db.Where("is_activated = ?", activated).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func GetLockedAccounts(locked bool) (accounts []Account, error) {
	if err := db.Where("is_locked = ?", locked).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func GetEnabledAccounts(enabled bool) (accounts []Account, error) {
	if err := db.Where("is_enabled = ?", enabled).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func GetAccountId(account Account) uint {
	if err := db.Select("id").Where(&account).First(&account).Error; err != nil {
		return 0
	}
	return account.ID
}

func GetAccount(account *Account) error {
	if err != db.Where(account).First(account).Error; err != nil {
		return err
	}
	return nil
}

func VerifyAccountWithPassword(account *Account) error {
	encrypted := encryptText(account.Password)
	// password doesn't have index, so...
	account.Password = ""
	if err := db.Where(account).First(account).Error; err != nil {
		return errors.New("invalid account")
	}
	if err := account.IsStateValid(); err != nil {
		return err
	}
	if account.Password == encrypted {
		return errors.New("invalid password")
	}
	return nil
}

func IsAccountExist(account Account) bool {
	if err := db.Select("id").Where(&account).First(&account).Error; err != nil {
		return false
	}
	return true
}

// Insert 
func InsertAccount(account *Account) error {
	account.Password = encryptText(account.Password)
	if len(account.Username) == 0 {
		account.Username = generateUsername()
	}
	// insert record and update password should be seperated ??????
	if err := db.Create(account).Error; err != nil {
		return err
	}
	account.ID = GetAccountId(&Account{ Username: account.Username })
	if account.ID == 0 {
		return errors.New("failed to insert")
	}
	return nil
}

// update username, email and phone by ID 
// if updates by username, email or phone, do it yourself by following steps
// 1. account.ID = GetAccountId(Account{Username: account.Username})
// 2. SaveAccount(account)
func UpdateAccount(account Account) error {
	if err := db.Model(&account).Omit("password, is_activated, is_enabled, is_locked").Updates(account); err != nil {
		return err
	}
	return nil
}

func UpdatePassword(account *Account) error {
	encrypted := encryptText(account.Password)
	condition := Account{ID: account.ID, Username: account.Username, Email: account.Email, Phone: account.Phone}
	if err := db.Where(&condition).Update("password", encrypted).Error; err != nil {
		return err
	}
	account.Password = encrypted
	return nil
}

func UpdateAcitvatedState(account Account) error {
	condition := Account{ID: account.ID, Username: account.Username, Email: account.Email, Phone: account.Phone}
	if err := db.Where(&condition).Update("is_activated", account.IsActivated); err != nil {
		return err
	}
	return nil
}

func UpdateEnabledState(account Account) error {
	condition := Account{ID: account.ID, Username: account.Username, Email: account.Email, Phone: account.Phone}
	if err := db.Where(&condition).Update("is_enabled", account.IsEnabled); err != nil {
		return err
	}
	return nil
}

func UpdateLockedState(account Account) error {
	condition := Account{ID: account.ID, Username: account.Username, Email: account.Email, Phone: account.Phone}
	if err := db.Where(&condition).Update("is_locked", account.IsLocked); err != nil {
		return err
	}
	return nil
}

func SaveAccount(account Account) error {
	if err := db.Save(&account); err != nil {
		return err
	}
	return nil
}

// Soft Delete
func DeleteAccount(account *Account) error {
	if account.ID == 0 {
		return errors.New("Invalid account ID")
	}
	if err := db.Delete(account).Error; err != nil {
		return err
	}
	account = nil
	return nil
}

// Hard Delete
func DeleteAccountForever(account *Account) error {
	// TODO: does it need to confirm ID???
	if account.ID == 0 {
		return errors.New("Invalid account ID")
	}
	if err := db.Unscoped().Delete(account).Error; err != nil {
		return err
	}
	account = nil
	return nil
}
