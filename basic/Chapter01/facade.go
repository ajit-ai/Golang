// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// FacadeAccount struct
type FacadeAccount struct {
	id          string
	accountType string
}

// FacadeAccount class method create - creates account given AccountType
func (account *FacadeAccount) create(accountType string) *FacadeAccount {
	fmt.Println("account creation with type")
	account.accountType = accountType

	return account
}

// FacadeAccount class method getById  given id string
func (account *FacadeAccount) getById(id string) *FacadeAccount {
	fmt.Println("getting account by Id")
	return account
}

// FacadeAccount class method deleteById given id string
func (account *FacadeAccount) deleteById(id string) {
	fmt.Println("delete account by id")
}

// FacadeCustomer struct
type FacadeCustomer struct {
	name string
	id   int
}

// FacadeCustomer class method create - create Customer given nam
func (customer *FacadeCustomer) create(name string) *FacadeCustomer {
	fmt.Println("creating customer")
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

// Transaction class method create Transaction
func (transaction *Transaction) create(srcAccountId string, destAccountId string, amount float32) *Transaction {
	fmt.Println("creating transaction")
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
	return transaction
}

// BranchManagerFacade struct
type BranchManagerFacade struct {
	account     *FacadeAccount
	customer    *FacadeCustomer
	transaction *Transaction
}

// methodd NewBranchManagerFacade
func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&FacadeAccount{}, &FacadeCustomer{}, &Transaction{}}
}

// BranchManagerFacade class method createCustomerAccount
func (facade *BranchManagerFacade) createCustomerAccount(customerName string, accountType string) (*FacadeCustomer, *FacadeAccount) {
	var customer = facade.customer.create(customerName)
	var account = facade.account.create(accountType)
	return customer, account
}

// BranchManagerFacade class method createTransaction
func (facade *BranchManagerFacade) createTransaction(srcAccountId string, destAccountId string, amount float32) *Transaction {

	var transaction = facade.transaction.create(srcAccountId, destAccountId, amount)
	return transaction

}

// FacadeMain method
func FacadeMain() {
	var facade = NewBranchManagerFacade()
	var customer *FacadeCustomer
	var account *FacadeAccount
	customer, account = facade.createCustomerAccount("Thomas Smith", "Savings")
	fmt.Println(customer.name)
	fmt.Println(account.accountType)
	var transaction = facade.createTransaction("21456", "87345", 1000)
	fmt.Println(transaction.amount)
}
