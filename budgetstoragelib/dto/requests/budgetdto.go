package requests

type BudgetRequest struct {
	User    string      `json:"user"`
	Budgets []BudgetDto `json:"budgets"`
}

type BudgetDto struct {
	Name   string  `json:"name"`
	Type   string  `json:"type,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	Remove bool    `json:"remove,omitempty"`
	Change BudgetChange `json:"change,omitempty"`
}

type BudgetChange struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}
