package models

import "time"

type RepeatableType int

const (
	Daily RepeatableType = iota
	Week
	Month
	Year
	Quater
	MonthLastDay
	FirstDayOfMonth
	PayPeriod
)

type BudgetItem struct {
	id             int
	name           string
	amount         float64
	repeatableType RepeatableType
	startsOn       time.Time
	endsOn				 *time.Time
	assignedBudget *Budget
}

func NewBudgetItem(id int, name string, amount float64, repeatableType RepeatableType, startsOn time.Time, endsOn time.Time, budget *Budget) *BudgetItem {
	return &BudgetItem{id, name, amount, repeatableType, startsOn, &endsOn, budget}
}

func (b BudgetItem) Name() string {
	return b.name
}

func (b BudgetItem) Amount() float64 {
	return b.amount
}

func (b BudgetItem) RepeatableType() RepeatableType {
	return b.repeatableType
}

func (b BudgetItem) StartsOn() time.Time {
	return b.startsOn
}

func (b BudgetItem) EndsOn() time.Time {
	if b.endsOn == nil {
		return time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)
	}

	return *b.endsOn
}

func (b BudgetItem) Budget() *Budget {
	return b.assignedBudget
}