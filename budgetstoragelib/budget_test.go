package budgetstoragelib_test

import (
	"budgetlib/models/budgettype"
	"budgetstoragelib/dto/requests"
	"budgetstoragelib/mssql"
	"budgetstoragelib/mssql/repositories"
	"testing"
)

func TestBudget(t *testing.T) {
	storage := mssql.GetInstance("server=localhost;user id=sa;password=h4Home@2017;port=1433;database=Budgets")

	if err := storage.Connect(); err != nil {
		t.Errorf("Expected storage.Connect to return nil, got %v", err)
	}

	defer storage.Close()

	budgetRepo := repositories.NewBudgetRepository(storage)

	body, err := budgetRepo.ManageBudgets(requests.BudgetRequest{
		User: "test",
		Budgets: []requests.BudgetDto{
			{
				Name:   "Test",
				Type:   "W",
				Amount: 100.0,
			},
		},
	})

	if err != nil {
		t.Errorf("Expected repo.ManageBudgets to return nil, got %v", err)
	}

	if len(body) == 0 {
		t.Errorf("Expected body to have a length greater than 0, got %v", len(body))
	} else if body[0].Name() != "Test" {
		t.Errorf("Expected body[0].Name() to return 'Test', got %v", body[0].Name())
	} else if body[0].Id() == 0 {
		t.Errorf("Expected body[0].Id() to return a value greater than 0, got %v", body[0].Id())
	} else if body[0].Amount() != 100.0 {
		t.Errorf("Expected body[0].Amount() to return 100.0, got %v", body[0].Amount())
	} else if body[0].Type() != budgettype.Weekly {
		t.Errorf("Expected body[0].Type() to return 'Weekly', got %v", body[0].Type())
	}

	insertedId := body[0].Id()

	body, err = budgetRepo.ManageBudgets(requests.BudgetRequest{
		User: "test",
		Budgets: []requests.BudgetDto{
			{
				Name: "Test",
				Change: requests.BudgetChange{
					Name:   "Test Budget",
					Amount: 200.0,
				},
			},
		},
	})

	if err != nil {
		t.Errorf("Expected repo.ManageBudgets to return nil, got %v", err)
	}

	if len(body) == 0 {
		t.Errorf("Expected body to have a length greater than 0, got %v", len(body))
	} else if body[0].Name() != "Test Budget" {
		t.Errorf("Expected body[0].Name() to return 'Test Budget', got %v", body[0].Name())
	} else if body[0].Id() != insertedId {
		t.Errorf("Expected body[0].Id() to return a value greater than 0, got %v", body[0].Id())
	} else if body[0].Amount() != 200.0 {
		t.Errorf("Expected body[0].Amount() to return 200.0, got %v", body[0].Amount())
	} else if body[0].Type() != budgettype.Weekly {
		t.Errorf("Expected body[0].Type() to return 'Weekly', got %v", body[0].Type())
	}


	body, err = budgetRepo.ManageBudgets(requests.BudgetRequest{
		User: "test",
		Budgets: []requests.BudgetDto{
			{
				Name:   "Test Budget",
				Remove: true,
			},
		},
	})

	if err != nil {
		t.Errorf("Expected repo.ManageBudgets to return nil, got %v", err)
	}

	if len(body) > 0 {
		t.Errorf("Expected body to have a length of 0, got %v", len(body))
	}
}
