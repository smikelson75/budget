package budgettype

type BudgetType int

const (
	Weekly BudgetType = iota
	BiWeekly	
	Monthly
	Unknown
)

func ConvertBudgetType(budgetType string) BudgetType {
	if budgetType == "W" {
		return Weekly
	} else if budgetType == "B" {
		return BiWeekly
	} else if budgetType == "M" {
		return Monthly
	} else {
		return Unknown
	}
}