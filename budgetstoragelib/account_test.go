package budgetstoragelib_test

import (
	"budgetlib/models"
	"testing"

	requests "budgetstoragelib/dto/requests"
	mssql "budgetstoragelib/mssql"
	repos "budgetstoragelib/mssql/repositories"
)

func TestAccount(t *testing.T) {
	storage := mssql.GetInstance("server=localhost;user id=sa;password=h4Home@2017;port=1433;database=Budgets")
	if err := storage.Connect(); err != nil {
		t.Errorf("Expected storage.Connect to return nil, got %v", err)
	}

	defer storage.Close()

	accountRepo := repos.NewAccountRepository(storage)

	body, err := accountRepo.ManageAccounts(requests.AccountRequest{
		User: "smikelson75@gmail.com",
		Accounts: []requests.AccountDto{
			{
				Name: "Test",
				AccountType: "C",
			},
		},
	})

	if err != nil {
		t.Errorf("Expected repo.ManageAccounts to return nil, got %v", err)
	}

	if len(body) == 0 {
		t.Errorf("Expected body to have a length greater than 0, got %v", len(body))
	} else if body[0].Name() != "Test" {
		t.Errorf("Expected body[0].Name() to return 'Test', got %v", body[0].Name())
	} else if body[0].Id() == 0 {
		t.Errorf("Expected body[0].Id() to return a value greater than 0, got %v", body[0].Id())
	}	else if body[0].Type() != models.Checking {
		t.Errorf("Expected body[0].AccountType() to return 'C', got %v", body[0].Type())
	}

	insertedId := body[0].Id()

	body, err = accountRepo.ManageAccounts(requests.AccountRequest{
		User: "smikelson75@gmail.com",
		Accounts: []requests.AccountDto{
			{
				Name: "Test",
				AccountType: "C",
				AcctChange: requests.AccountChange{
					Name: "Test Checking",
				},
			},
		},
	})

	if err != nil {
		t.Errorf("Expected repo.ManageAccounts to return nil, got %v", err)
	}

	if len(body) == 0 {
		t.Errorf("Expected body to have a length greater than 0, got %v", len(body))
	} else if body[0].Name() != "Test Checking" {
		t.Errorf("Expected body[0].Name() to return 'Test Checking', got %v", body[0].Name())
	} else if body[0].Id() != insertedId {
		t.Errorf("Expected body[0].Id() to return a value greater than 0, got %v", body[0].Id())
	} else if body[0].Type() != models.Checking {
		t.Errorf("Expected body[0].AccountType() to return 'C', got %v", body[0].Type())
	}

	body, err = accountRepo.ManageAccounts(requests.AccountRequest{
		User: "smikelson75@gmail.com",
		Accounts: []requests.AccountDto{
			{
				Name: "Test Checking",
				Remove: true,
			},
		},
	})

	if err != nil {
		t.Errorf("Expected repo.ManageAccounts to return nil, got %v", err)
	}

	if len(body) > 0 {
		t.Errorf("Expected body to have a length equal to 0, got %v", len(body))
	}
}
