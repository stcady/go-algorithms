package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// AccountDetails struct
type AccountDetails struct {
	id          string
	accountType string
}

// Account struct
type Account struct {
	details      *AccountDetails
	CustomerName string
}

// Account class method setDetails
func (account *Account) setDetails(id string, accountType string) {
	account.details = &AccountDetails{id, accountType}
}

// Account class method getId
func (account Account) getId() string {
	return account.details.id
}

// Account class method getAccountType
func (account Account) getAccountType() string {
	return account.details.accountType
}

func main() {
	account := &Account{CustomerName: "John Smith"}
	account.setDetails("4532", "current")
	jsonAccount, err := json.Marshal(account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Private Class Hidden: ", string(jsonAccount))
	fmt.Println("Account ID: ", account.getId())
	fmt.Println("Account Type: ", account.getAccountType())
}
