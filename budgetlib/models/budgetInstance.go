package models

import "time"

type BudgetInstance struct {
	name string
	amount float64
	startDate time.Time
	endDate time.Time
}

func NewBugetInstance(name string, amount float64, startDate time.Time, endDate time.Time) *BudgetInstance {
	return &BudgetInstance{name, amount, startDate, endDate}
}

func (b BudgetInstance) Name() string {
	return b.name
}

func (b BudgetInstance) Amount() float64 {
	return b.amount
}

func (b BudgetInstance) StartDate() time.Time {
	return b.startDate
}

func (b BudgetInstance) EndDate() time.Time {
	return b.endDate
}