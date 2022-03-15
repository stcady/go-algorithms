package main

import "fmt"

// Account struct
type Account struct {
	id          string
	accountType string
}

// create method for the Account class
func (account *Account) create(accountType string) *Account {
	fmt.Println("account creation with type")
	account.accountType = accountType
	return account
}

// deleteById method for the Account class
func (account *Account) deleteById(id string) {
	fmt.Println("delete account by id")
}

// Customer struct
type Customer struct {
	name string
	id   string
}

// create method for the Customer class
func (customer *Customer) create(name string) *Customer {
	fmt.Println("create customer")
	customer.name = name
	return customer
}

// Transaction struct
type Transaction struct {
	id            string
	amount        float32
	srcAccountId  string
	destAccountId string
}

// create method for Transaction class
func (transaction *Transaction) create(srcAccountId string, destAccountId string, amount float32) *Transaction {
	fmt.Println("creating transaction")
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
	return transaction
}

// BranchManagerFacade struct
type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

// NewBranchManagerFacade function
func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

// createCustomerAccount method for the BranchManagerFacade class
func (facade *BranchManagerFacade) createCustomerAccount(customerName string, accountType string) (*Customer, *Account) {
	customer := facade.customer.create(customerName)
	account := facade.account.create(accountType)
	return customer, account
}

// createTransaction method for the BranchManagerFacade class
func (facade *BranchManagerFacade) createTransaction(srcAccountId string, destAccountId string, amount float32) *Transaction {
	transaction := facade.transaction.create(srcAccountId, destAccountId, amount)
	return transaction
}

func main() {
	facade := NewBranchManagerFacade()
	customer, account := facade.createCustomerAccount("Thomas Smith", "Savings")
	fmt.Println(facade.customer.name)
	fmt.Println(customer.name)
	fmt.Println(facade.account.accountType)
	fmt.Println(account.accountType)
	transaction := facade.createTransaction("21456", "87345", 1000)
	fmt.Println(facade.transaction.amount)
	fmt.Println(transaction.amount)
}
