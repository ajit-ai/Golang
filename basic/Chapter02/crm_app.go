// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,database/sql, net/http, text/template package
import (
	"fmt"
	"net/http"
	"text/template"
	//    "errors"
	"log"
)

var template_html = mustParseTemplates()

// mustParseTemplates loads the CRM templates; it falls back to an empty
// template set when the templates directory is not present in the working directory
func mustParseTemplates() *template.Template {
	tmpl, err := template.ParseGlob(`templates/*`)
	if err != nil {
		return template.New(`none`)
	}
	return tmpl
}

func Home(writer http.ResponseWriter, request *http.Request) {
	var customers []CrmCustomer
	customers = GetCrmCustomers()
	log.Println(customers)
	template_html.ExecuteTemplate(writer, "Home", customers)

}

func Create(writer http.ResponseWriter, request *http.Request) {
	//  var customers []CrmCustomer
	//customers = GetCrmCustomers()
	//log.Println(customers)
	template_html.ExecuteTemplate(writer, "Create", nil)

}

func Insert(writer http.ResponseWriter, request *http.Request) {
	//  var customers []CrmCustomer
	//customers = GetCrmCustomers()
	//log.Println(customers)
	var customer CrmCustomer
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")
	InsertCrmCustomer(customer)
	var customers []CrmCustomer
	customers = GetCrmCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)

}
func Alter(writer http.ResponseWriter, request *http.Request) {
	//  var customers []CrmCustomer
	//customers = GetCrmCustomers()
	//log.Println(customers)
	var customer CrmCustomer
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer.CustomerId = customerId
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")
	UpdateCrmCustomer(customer)
	var customers []CrmCustomer
	customers = GetCrmCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)

}

func Update(writer http.ResponseWriter, request *http.Request) {

	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer CrmCustomer
	customer = GetCrmCustomerById(customerId)
	//log.Println(customer)
	//var customers []CrmCustomer
	//customers = GetCrmCustomers()
	//log.Println(customers)
	template_html.ExecuteTemplate(writer, "Update", customer)

}

func Delete(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer CrmCustomer
	customer = GetCrmCustomerById(customerId)
	DeleteCrmCustomer(customer)
	var customers []CrmCustomer
	customers = GetCrmCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)

}

func View(writer http.ResponseWriter, request *http.Request) {
	//var customers []CrmCustomer
	//customers = GetCrmCustomers()
	//log.Println(customers)
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer CrmCustomer
	customer = GetCrmCustomerById(customerId)
	fmt.Println(customer)
	var customers []CrmCustomer
	customers = []CrmCustomer{customer}
	//  customers.append(customer)
	template_html.ExecuteTemplate(writer, "View", customers)

}

func CrmAppMain() {
	log.Println("Server started on: http://localhost:8000")
	//  var template_html *template.Template
	//template_html = template.Must(template.ParseFiles("main.html"))
	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8000", nil)
}
