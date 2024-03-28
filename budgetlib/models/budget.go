package models

import (
	"budgetlib/models/budgettype"
	"time"
)

type Budget struct {
	id         int
	name       string
	budgetType budgettype.BudgetType
	amount     float64
}

func NewBudget(id int, name string, budgetType budgettype.BudgetType, amount float64) *Budget {
	return &Budget{id, name, budgetType, amount}
}

func (b Budget) Instantiate(startDate time.Time) *BudgetInstance {
	if b.budgetType == budgettype.Weekly {
		return NewBugetInstance(b.name, b.amount, startDate, startDate.AddDate(0, 0, 7))
	} else if b.budgetType == budgettype.BiWeekly {
		return NewBugetInstance(b.name, b.amount, startDate, startDate.AddDate(0, 0, 14))
	} else if b.budgetType == budgettype.Monthly {
		return NewBugetInstance(b.name, b.amount, startDate, startDate.AddDate(0, 1, 0))
	}

	return nil
}

func (b Budget) Id() int {
	return b.id
}

func (b Budget) Name() string {
	return b.name
}

func (b Budget) Type() budgettype.BudgetType {
	return b.budgetType
}

func (b Budget) Amount() float64 {
	return b.amount
}
