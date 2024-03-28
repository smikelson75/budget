package models

import "time"

type Expense struct {
	id int
	name string
	amount float64
	estimatedDueDate time.Time
	instance *BudgetInstance
}

func NewExpense(id int, name string, amount float64, estimatedDueDate time.Time, instance *BudgetInstance) *Expense {
	return &Expense{id, name, amount, estimatedDueDate, instance}
}

func (e Expense) Name() string {
	return e.name
}

func (e Expense) Amount() float64 {
	return e.amount
}

func (e Expense) EstimatedDueDate() time.Time {
	return e.estimatedDueDate
}

func (e Expense) Instance() *BudgetInstance {
	return e.instance
}
