// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,database/sql, net/http, text/template package
import (
	//    "fmt"
	"database/sql"
	//  "log"
	//  "net/http"
	//  "text/template"
	//  "errors"
	_ "github.com/go-sql-driver/mysql"
)

type CrmCustomer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

func GetCrmConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "newuser"
	databasePass := "newuser"
	databaseName := "crm"
	database, error := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}

func GetCrmCustomerById(customerId int) CrmCustomer {
	var database *sql.DB
	database = GetCrmConnection()

	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM CrmCustomer WHERE CustomerId=?", customerId)
	if error != nil {
		panic(error.Error())
	}
	//fmt.Println(rows)
	var customer CrmCustomer

	for rows.Next() {
		var customerId int
		var customerName string
		var SSN string
		error = rows.Scan(&customerId, &customerName, &SSN)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = SSN
	}
	defer database.Close()
	return customer
}

func GetCrmCustomers() []CrmCustomer {
	var database *sql.DB
	database = GetCrmConnection()

	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM CrmCustomer ORDER BY Customerid DESC")
	if error != nil {
		panic(error.Error())
	}
	var customer CrmCustomer

	var customers []CrmCustomer
	customers = []CrmCustomer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}

	defer database.Close()

	return customers
}

func InsertCrmCustomer(customer CrmCustomer) {
	var database *sql.DB
	database = GetCrmConnection()

	var error error
	var insert *sql.Stmt
	insert, error = database.Prepare("INSERT INTO CrmCustomer(CustomerName,SSN) VALUES(?,?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)
	//log.Println("INSERT: CrmCustomer Name: " + CrmCustomer.name + " | SSN: " + CrmCustomer.ssn)

	defer database.Close()

	//return CrmCustomer{}
}

func UpdateCrmCustomer(customer CrmCustomer) {
	var database *sql.DB
	database = GetCrmConnection()

	var error error
	var update *sql.Stmt
	update, error = database.Prepare("UPDATE CrmCustomer SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if error != nil {
		panic(error.Error())
	}
	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	//log.Println("INSERT: CrmCustomer Name: " + CrmCustomer.name + " | SSN: " + CrmCustomer.ssn)

	defer database.Close()

	//return CrmCustomer{}
}
func DeleteCrmCustomer(customer CrmCustomer) {
	var database *sql.DB
	database = GetCrmConnection()

	var error error
	var delete *sql.Stmt
	delete, error = database.Prepare("DELETE FROM CrmCustomer WHERE Customerid=?")
	if error != nil {
		panic(error.Error())
	}
	delete.Exec(customer.CustomerId)
	//log.Println("INSERT: CrmCustomer Name: " + CrmCustomer.name + " | SSN: " + CrmCustomer.ssn)

	defer database.Close()

	//return CrmCustomer{}
}

/*func main() {

     var customers []CrmCustomer
    customers = GetCrmCustomers()
    fmt.Println(customers)

  //  var CrmCustomer CrmCustomer
//    CrmCustomer.CustomerName = "Thomas Smith"
  //  CrmCustomer.SSN = "2323343"

  //  InsertCrmCustomer(CrmCustomer)

  //var CrmCustomer CrmCustomer
  //  CrmCustomer.CustomerName = "George Thompson"
  //  CrmCustomer.SSN = "23233432"
  //  CrmCustomer.CustomerId = 2
  //  UpdateCrmCustomer(CrmCustomer)

var CrmCustomer CrmCustomer
  //CrmCustomer.CustomerName = "George Thompson"
  //CrmCustomer.SSN = "23233432"
 customer.CustomerId = 2

    DeleteCrmCustomer(CrmCustomer)
    customers = GetCrmCustomers()
    fmt.Println(customers)


}
*/
