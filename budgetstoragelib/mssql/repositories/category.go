package repositories

import (
	"budgetlib/models"
	"budgetstoragelib/dto/requests"
	"budgetstoragelib/dto/responses"
	"budgetstoragelib/interfaces"
	"budgetstoragelib/liberrors"
	"database/sql"
	"encoding/json"
	"fmt"
)

type CategoryRepository struct {
	server interfaces.IServer
	factory liberrors.Factory
}

func NewCategoryRepository(server interfaces.IServer) *CategoryRepository {
	return &CategoryRepository{server, liberrors.Factory{}}
}

func (r CategoryRepository) Manage(input requests.CategoryRequest) ([]*models.Category, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, r.factory.WrapError(err, "category")
	}

	var output string

	db := r.server.Database()
	_, err = db.Exec(
		"[Org].[usp_ManageCategories]",
		sql.Named("jsonDoc", string(payload)),
		sql.Named("outputJson", sql.Out{Dest: &output}),
	)
	if err != nil {
		return nil, r.factory.WrapError(err, "category")
	}

	var categories responses.CategoryResponse = responses.CategoryResponse{}
	err = json.Unmarshal([]byte(output), &categories)
	if err != nil {
		return nil, r.factory.WrapError(err, "category")
	}

	var results []*models.Category = []*models.Category{}
	for key, category := range categories.Categories {
		results = append(results, &models.Category{Id: category.Id, Name: category.Name})
		fmt.Printf("results: %v\n", results[key])

	}

	return results, nil
}

func (r CategoryRepository) ListCategories() ([]*models.Category, error) {
	var categories responses.CategoryResponse = responses.CategoryResponse{}
	var output string

	db := r.server.Database()
	row := db.QueryRow("SELECT [Org].[udf_ListCategories]()")

	err := row.Scan(&output)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		return []*models.Category{}, nil
	}

	err = json.Unmarshal([]byte(output), &categories)
	if err != nil {
		return nil, r.factory.WrapError(err, "category")
	}

	var results []*models.Category = []*models.Category{}
	for key, category := range categories.Categories {
		results = append(results, &models.Category{ Id: category.Id, Name: category.Name})
		fmt.Printf("results: %v\n", results[key])

	}

	return results, nil
}
