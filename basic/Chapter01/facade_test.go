package main

import "testing"

func TestNewBranchManagerFacade(t *testing.T) {
	facade := NewBranchManagerFacade()
	if facade.account == nil {
		t.Error("facade.account is nil, want initialized")
	}
	if facade.customer == nil {
		t.Error("facade.customer is nil, want initialized")
	}
	if facade.transaction == nil {
		t.Error("facade.transaction is nil, want initialized")
	}
}

func TestCreateCustomerAccount(t *testing.T) {
	facade := NewBranchManagerFacade()
	customer, account := facade.createCustomerAccount("Thomas Smith", "Savings")
	if customer == nil || customer.name != "Thomas Smith" {
		t.Errorf("customer.name = %q, want %q", customer.name, "Thomas Smith")
	}
	if account == nil || account.accountType != "Savings" {
		t.Errorf("account.accountType = %q, want %q", account.accountType, "Savings")
	}
}

func TestCreateTransaction(t *testing.T) {
	facade := NewBranchManagerFacade()
	transaction := facade.createTransaction("21456", "87345", 1000)
	if transaction.srcAccountId != "21456" {
		t.Errorf("srcAccountId = %q, want %q", transaction.srcAccountId, "21456")
	}
	if transaction.destAccountId != "87345" {
		t.Errorf("destAccountId = %q, want %q", transaction.destAccountId, "87345")
	}
	if transaction.amount != 1000 {
		t.Errorf("amount = %v, want 1000", transaction.amount)
	}
}

func TestFacadeMain(t *testing.T) {
	FacadeMain()
}
