package model

import (
	"fmt"
	"testing"
)

func init() {
	// drop account_test table
	if err := DropTable("accounts"); err != nil {
		fmt.Printf("drop table error: %v\n", err)
	}
	// migrate new account_table
	db.AutoMigrate(&Account{})
}

func TestInsertAccount(t *testing.T) {
	account1 := Account{Username: "Ai3ueJSKt6", Email: "omg@gmail.com", Phone: "13888888888", Password: "88888888", IsActivated: true, IsEnabled: true, IsLocked: false}
	if err := InsertAccount(&account1); err != nil {
		t.Errorf("insert account1 error: %v", err)
	}
	account2 := Account{Email: "omg2@gmail.com", Password: "88888888", IsActivated: false, IsEnabled: false, IsLocked: false}
	if err := InsertAccount(&account2); err != nil {
		t.Errorf("insert account2 error: %v", err)
	}
	account3 := Account{Phone: "13999999999", Password: "88888888", IsActivated: false, IsEnabled: false, IsLocked: false}
	if err := SaveAccount(&account3); err != nil {
		t.Errorf("insert account3 error: %v", err)
	}
	account4 := Account{Phone: "13777777777", Password: "88888888", IsActivated: false, IsEnabled: false, IsLocked: false}
	if err := SaveAccount(&account4); err != nil {
		t.Errorf("insert account4 error: %v", err)
	}
}

func TestQueryAllAccounts(t *testing.T) {
	accounts, err := GetAccounts()
	if err != nil {
		t.Errorf("query all accounts error: %v", err)
	}
	if len(accounts) != 4 {
		t.Errorf("unexpected amount of accounts")
	}
	if accounts[0].Username != "Ai3ueJSKt6" {
		t.Errorf("unexpected username: %v", accounts[0])
	}
	if accounts[1].Email != "omg2@gmail.com" {
		t.Errorf("unexpected email: %v", accounts[1])
	}
	if accounts[2].Phone != "13999999999" {
		t.Errorf("unexpected phone number: %v", accounts[2])
	}
	if accounts[3].Phone != "13777777777" {
		t.Errorf("unexpected phone number: %v", accounts[3])
	}
}

func TestQueryValidAccounts(t *testing.T) {
	accounts, err := GetValidAccounts()
	if err != nil {
		t.Errorf("query valid accounts error: %v", err)
	}
	if len(accounts) != 1 {
		t.Errorf("unexpected amount of accounts")
	}
}

func TestQueryActivatedAccounts(t *testing.T) {
	accounts, err := GetActivatedAccounts(true)
	if err != nil {
		t.Errorf("query activated accounts error: %v", err)
	}
	if len(accounts) != 1 {
		t.Errorf("unexpected amount of accounts")
	}
}

func TestQueryAccountId(t *testing.T) {
	accountId := GetAccountId(Account{Username: "Ai3ueJSKt6"})
	if accountId == 0 {
		t.Errorf("query account id error")
	}
	if accountId != 1 {
		t.Errorf("unexpected account id, expected: %d, actual: %d", 1, accountId)
	}
}

func TestQueryAccount(t *testing.T) {
	account, err := GetAccount(Account{Username: "Ai3ueJSKt6"})
	if err != nil {
		t.Errorf("query account id error: %v", err)
	}
	if account.ID != 1 {
		t.Errorf("unexpected account id, expected: %d, actual: %d", 1, account.ID)
	}
}

func TestVerifyAccountWithPassword(t *testing.T) {
	account, _ := VerifyAccountWithPassword(Account{Username: "Ai3ueJSKt6", Password: "88888888"})
	if account == nil {
		t.Errorf("unmatched account(%s) and password(%s)", account.Username, account.Password)
	}
	account2, _ := VerifyAccountWithPassword(Account{Phone: "13777777777", Password: "88888888"})
	if account2 != nil {
		t.Errorf("verification should be fail, but the actual is not")
	}
}

func TestIsAccountExist(t *testing.T) {
	if !IsAccountExist(Account{Username: "Ai3ueJSKt6"}) {
		t.Errorf("account [%s] should be exist", "Ai3ueJSKt6")
	}
	if IsAccountExist(Account{Phone: "1111"}) {
		t.Errorf("account [1111] should not be exist")
	}
}

func TestUpdateAccount(t *testing.T) {
	account, err := GetAccount(Account{Phone: "13777777777"})
	if err != nil {
		t.Errorf("account query error: %v", err)
		return
	}
	account.Email = "123@gmail.com"
	account.Username = "1234567890"
	account.IsActivated = true
	if err = UpdateAccount(*account); err != nil {
		t.Errorf("update account error: %v", err)
	}
	account.IsActivated = false
	if err = UpdateAccount(*account); err != nil {
		t.Errorf("update account error: %v", err)
	}
}

func TestUpdatePassword(t *testing.T) {
	if err := UpdatePassword(Account{Email: "omg2@gmail.com", Password: "888888"}); err != nil {
		t.Errorf("update password error: %v", err)
	}
	if err := UpdateActivatedState(Account{Email: "omg2@gmail.com", IsActivated: true}); err != nil {
		t.Errorf("update activated state error: %v", err)
	}
	if err := UpdateEnabledState(Account{Email: "omg2@gmail.com", IsEnabled: true}); err != nil {
		t.Errorf("update enabled state error: %v", err)
	}
	if err := UpdateLockedState(Account{Email: "omg2@gmail.com", IsLocked: false}); err != nil {
		t.Errorf("update locked state error: %v", err)
	}
	if account, err := VerifyAccountWithPassword(Account{Email: "omg2@gmail.com", Password: "88888888"}); account != nil {
		t.Errorf("password has been updated, shouldn't be matched, error: %v", err)
	}
	if account, err := VerifyAccountWithPassword(Account{Email: "omg2@gmail.com", Password: "888888"}); account == nil {
		t.Errorf("new password should be matched, error: %v", err)
	}
}

func TestDelete(t *testing.T) {
	account, err := GetAccount(Account{Email: "omg2@gmail.com"})
	if err != nil {
		t.Errorf("get account error: %v", err)
	}
	if err = DeleteAccount(*account); err != nil {
		t.Errorf("delete error: %v", err)
	}
}

func TestDeleteForever(t *testing.T) {
	account, err := GetAccount(Account{Phone: "13999999999"})
	if err != nil {
		t.Errorf("get account error: %v", err)
	}
	if err = DeleteAccountForever(*account); err != nil {
		t.Errorf("delete error: %v", err)
	}
}
