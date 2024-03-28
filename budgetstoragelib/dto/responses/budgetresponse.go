package responses

type BudgetResponse struct {
	Budgets []BudgetInternal `json:"budgets"`
}

type BudgetInternal struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}
