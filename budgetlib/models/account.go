package models

import "strings"

type AccountType int

const (
	Checking AccountType = iota
	Savings
	Credit
	Loan
	Other
	Unknown
)

type Account struct {
	id int
	name string
	accountType AccountType
}

func NewAccount(id int, name string, accountType AccountType) *Account {
	return &Account{id, name, accountType}
}

func (a Account) Name() string {
	return a.name
}

func (a Account) Type() AccountType {
	return a.accountType
}

func (a Account) Id() int {
	return a.id
}

func ConvertAccountType(accountType string) AccountType {
	switch strings.ToUpper(accountType) {
	case "C":
		return Checking
	case "S":
		return Savings
	case "R":
		return Credit
	case "L":
		return Loan
	case "O":
		return Other
	default:
		return Unknown
	}
}