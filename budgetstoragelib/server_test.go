package budgetstoragelib_test

import (
	requests "budgetstoragelib/dto/requests"
	mssql "budgetstoragelib/mssql"
	repos "budgetstoragelib/mssql/repositories"
	"fmt"
	"testing"

	_ "github.com/microsoft/go-mssqldb"
)

func TestCategoryRepository(t *testing.T) {
	storage := mssql.GetInstance("server=localhost;user id=sa;password=h4Home@2017;port=1433;database=Budgets")

	if err := storage.Connect(); err != nil {
		t.Errorf("Expected storage.Connect to return nil, got %v", err)
	}

	defer storage.Close()

	categoryRepo := repos.NewCategoryRepository(storage)

	body, err := categoryRepo.Manage(requests.CategoryRequest{
		User: "smikelson75@gmail.com",
		Categories: []requests.CategoryDto{
			{
				Name: "Personal",
			},
		},
	})

	if err != nil {
		t.Errorf("Expected repo.Manage to return nil, got %v", err)
	}

	if len(body) == 0 {
		t.Errorf("Expected body to have a length greater than 0, got %v", len(body))
	} else if body[0].Name != "Personal" {
		t.Errorf("Expected body[0].Name() to return 'Personal', got %v", body[0].Name)
	} else if body[0].Id == 0 {
		t.Errorf("Expected body[0].Id() to return a value greater than 0, got %v", body[0].Id)
	}

	insertedId := body[0].Id

	body, err = categoryRepo.Manage(requests.CategoryRequest{
		User: "smikelson75@gmail.com",
		Categories: []requests.CategoryDto{
			{
				Id:   insertedId,
				Name: "Personal Expenses",
			},
		},
	})

	if err != nil {
		t.Errorf("Expected repo.Manage to return nil, got %v", err)
	}

	if len(body) == 0 {
		t.Errorf("Expected body to have a length greater than 0, got %v", len(body))
	} else if body[0].Name != "Personal Expenses" {
		t.Errorf("Expected body[0].Name() to return 'Personal Expenses', got %v", body[0].Name)
	} else if body[0].Id != insertedId {
		t.Errorf("Expected body[0].Id() to return a value greater than 0, got %v", body[0].Id)
	}

	body, err = categoryRepo.Manage(requests.CategoryRequest{
		User: "smikelson75@gmail.com",
		Categories: []requests.CategoryDto{
			{
				Id:     insertedId,
				Remove: true,
			},
		},
	})

	if err != nil {
		t.Errorf("Expected repo.Manage to return nil, got %v", err)
	}

	fmt.Printf(`Output: %v\n`, body)
	if len(body) > 0 {
		t.Errorf("Expected body to have a length equal to 0, got %v", len(body))
	}
}

func TestCategoryRepositoryListCategories(t *testing.T) {
	storage := mssql.GetInstance("server=localhost;user id=sa;password=h4Home@2017;port=1433;database=Budgets")

	if err := storage.Connect(); err != nil {
		t.Errorf("Expected storage.Connect to return nil, got %v", err)
	}

	defer storage.Close()

	categoryRepo := repos.NewCategoryRepository(storage)

	body, err := categoryRepo.ListCategories()

	if err != nil {
		t.Errorf("Expected repo.ListCategories to return nil, got %v", err)
	}

	if len(body) > 0 {
		t.Errorf("Expected body to have a length greater than 0, got %v", len(body))
	}

	body, err = categoryRepo.Manage(requests.CategoryRequest{
		User: "smikelson75@gmail.com",
		Categories: []requests.CategoryDto{
			{
				Name: "Personal Expenses",
			},
			{
				Name: "Business Expenses",
			},
		},
	})

	if err != nil {
		t.Errorf("Expected repo.Manage to return nil, got %v", err)
	} else if len(body) != 2 {
		t.Errorf("Expected body to have a length equal to 2, got %v", len(body))
	}

	body, err = categoryRepo.ListCategories()

	if err != nil {
		t.Errorf("Expected repo.ListCategories to return nil, got %v", err)
	}

	request := requests.CategoryRequest{
		User: "smikelson75@gmail.com",
	}

	for _, v := range body {
		request.Categories = append(request.Categories, requests.CategoryDto{
			Id: v.Id,
			Remove: true,
		})
	}

	body, err = categoryRepo.Manage(request)

	if err != nil {
		t.Errorf("Expected repo.Manage to return nil, got %v", err)
	} else if len(body) > 0 {
		t.Errorf("Expected body to have a length equal to 0, got %v", len(body))
	}

}
