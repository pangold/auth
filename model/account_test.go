package main

import (
	"fmt"

	"./model"
)

func testAutoSave() {
	account := model.Account{Username: "Ai3ueJSKt6", Email: "omg@gmail.com", Phone: "13888888888", Password: "88888888", IsActivated: false, IsEnabled: false, IsLocked: false}
	if err := model.SaveAccount(&account); err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
	account.Email = "god@gmail.com"
	fmt.Println(account)
	if err := model.SaveAccount(&account); err != nil {
		fmt.Println(err)
	}
}

func testQuery() {
	var err error
	var accounts, accounts2, accounts3, accounts4 []model.Account
	accounts, err = model.GetAccounts()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, account := range accounts {
		fmt.Println(account)
	}
	//
	accounts2, err = model.GetActivatedAccounts(false)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, account := range accounts2 {
		fmt.Println(account)
	}
	//
	accounts3, err = model.GetEnabledAccounts(false)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, account := range accounts3 {
		fmt.Println(account)
	}
	//
	accounts4, err = model.GetLockedAccounts(false)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, account := range accounts4 {
		fmt.Println(account)
	}
	//
	account, _ := model.GetAccountByPhone("13888888888")
	fmt.Println(account)
	// 
	account2, _ := model.GetAccountByUsername("Ai3ueJSKt6")
	fmt.Println(account2)
	//
	account3, _ := model.GetAccountByEmail("god@gmail.com")
	fmt.Println(account3)
	//
	var id1, id2, id3 uint
	id1 = model.GetAccountIdByPhone("13888888888")
	fmt.Println(id1)
	id2 = model.GetAccountIdByUsername("Ai3ueJSKt6")
	fmt.Println(id2)
	id3 = model.GetAccountIdByEmail("god@gmail.com")
	fmt.Println(id3)
}

func testPasswordMatch() {
	// With correct/incorrect password
	var err error
	var account10, account11, account12 *model.Account
	if account10, err = model.MatchPhoneAndPassword("13888888888", "88888888"); err != nil {
		fmt.Println(err.Error())
	}
	if account11, err = model.MatchUsernameAndPassword("Ai3ueJSKt6", "88888888"); err != nil {
		fmt.Println(err.Error())
	}
	if account12, err = model.MatchEmailAndPassword("god@gmail.com", "8888888"); err != nil { // incorrect
		fmt.Println(err.Error())
	}
	fmt.Println("MatchXXXAndPassword Begin")
	fmt.Println(account10)
	fmt.Println(account11)
	fmt.Println(account12)
	fmt.Println("MatchXXXAndPassword End")
}

func testInsert() {
	fmt.Println("TestInsert Begin")
	account := model.Account{Username: "1234567890", Password: "88888888"}
	if err := model.InsertAccount(&account); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(account)
	fmt.Println("TestInsert End")
}

func testInsertAndUpdate() {
	fmt.Println("TestInsertAndUpdate Begin")
	account := model.Account{Email: "x@111.com", Phone: "13111111111", Username: "pandora", Password: "12345678"}
	if err := model.InsertAccount(&account); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(account)
	account.Email = "pangold@163.com"
	account.Phone = "18828883888"
	if err := model.UpdateAccount(&account); err != nil {
		fmt.Println(err.Error())
	}
	//
	account2, _ := model.GetAccountByEmail("pangold@163.com")
	fmt.Println(account2)
	fmt.Println("TestInsertAndUpdate End")
}

func testUpdateStatus() {
	var err error
	var account model.Account
	if account, err = model.GetAccountByUsername("1234567890"); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(account)
	//
	account.Email = "abcd@gmail.com"
	account.Phone = "13538283577"
	if err = model.SaveAccount(&account); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(account)
	//
	if err = model.EnableAccount(&account); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(account)
	//
	if err = model.ActivateAccount(&account); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(account)
	//
	if err = model.LockAccount(&account); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(account)
	//
	//if err = model.DeleteAccount(&account); err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(account)
}

func main() {
	//testAutoSave()
	//testQuery()
	//testPasswordMatch()
	//testInsert()
	//testInsertAndUpdate()
	testUpdateStatus()
}
