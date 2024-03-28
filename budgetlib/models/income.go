package models

type Income struct {
	id int
	name string
	amount float64
	instance *BudgetInstance
}

func NewIncome(id int, name string, amount float64, instance *BudgetInstance) *Income {
	return &Income{id, name, amount, instance}
}

func (i Income) Name() string {
	return i.name
}

func (i Income) Amount() float64 {
	return i.amount
}

func (i Income) Instance() *BudgetInstance {
	return i.instance
}