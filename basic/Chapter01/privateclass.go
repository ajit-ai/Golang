// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and encoding/json pacakges
import (
	"encoding/json"
	"fmt"
)

// AccountDetails struct
type AccountDetails struct {
	id          string
	accountType string
}

// PrivateClassAccount struct
type PrivateClassAccount struct {
	details      *AccountDetails
	CustomerName string
}

// PrivateClassAccount class method setDetails
func (account *PrivateClassAccount) setDetails(id string, accountType string) {

	account.details = &AccountDetails{id, accountType}
}

// PrivateClassAccount class method getId
func (account *PrivateClassAccount) getId() string {

	return account.details.id
}

// PrivateClassAccount class method getAccountType
func (account *PrivateClassAccount) getAccountType() string {

	return account.details.accountType
}

// PrivateClassMain method
func PrivateClassMain() {

	var account *PrivateClassAccount = &PrivateClassAccount{CustomerName: "John Smith"}
	account.setDetails("4532", "current")

	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Private Class hidden", string(jsonAccount))

	fmt.Println("Account Id", account.getId())

	fmt.Println("Account Type", account.getAccountType())

}
