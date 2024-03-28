package repositories

import (
	"budgetlib/models"
	"budgetlib/models/budgettype"
	"budgetstoragelib/dto/requests"
	"budgetstoragelib/dto/responses"
	"budgetstoragelib/interfaces"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

type BudgetRepostiory struct {
	server interfaces.IServer
}

func NewBudgetRepository(server interfaces.IServer) *BudgetRepostiory {
	return &BudgetRepostiory{server}
}

func (r BudgetRepostiory) ManageBudgets(input requests.BudgetRequest) ([]*models.Budget, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, errors.New("failed to marshal input for call to usp_ManageBudgets")
	}

	var output string

	db := r.server.Database()
	_, err = db.Exec(
		"[Budget].[usp_ManageBudgets]",
		sql.Named("jsonDoc", string(payload)),
		sql.Named("outputJson", sql.Out{Dest: &output}),
	)

	if err != nil {
		return nil, err
	}

	var budgets responses.BudgetResponse = responses.BudgetResponse{}
	err = json.Unmarshal([]byte(output), &budgets)
	if err != nil {
		return nil, err
	}

	var results []*models.Budget = []*models.Budget{}
	for key, budget := range budgets.Budgets {
		results = append(
			results,
			models.NewBudget(
				budget.Id,
				budget.Name,
				budgettype.ConvertBudgetType(budget.Type),
				budget.Amount,
			),
		)
		fmt.Printf("results: %v\n", results[key])

	}

	return results, nil
}
