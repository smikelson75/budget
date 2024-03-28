package budgetlib_test

import (
	"testing"
	"time"

	model "budgetlib/models"
	"budgetlib/models/budgettype"
)

func TestBudgetInstance(t *testing.T) {
	budget := model.NewBudget(-1, "Test Budget", budgettype.Weekly, 100.0)
	startDate := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	budgetInstance := budget.Instantiate(startDate)
	if budgetInstance.Name() != "Test Budget" {
		t.Errorf("Expected budgetInstance.Name to be 'Test Budget', got %s", budgetInstance.Name())
	}
	if budgetInstance.Amount() != 100.0 {
		t.Errorf("Expected budgetInstance.Amount to be 100.0, got %f", budgetInstance.Amount())
	}
	if budgetInstance.StartDate() != startDate {
		t.Errorf("Expected budgetInstance.StartDate to be %s, got %s", startDate, budgetInstance.StartDate())
	}
	expectedEndDate := time.Date(2019, time.January, 8, 0, 0, 0, 0, time.UTC)
	if budgetInstance.EndDate() != expectedEndDate {
		t.Errorf("Expected budgetInstance.EndDate to be %s, got %s", expectedEndDate, budgetInstance.EndDate())
	}
}

func TestBudgetMonthly(t *testing.T) {
	budget := model.NewBudget(-1, "Test Budget", budgettype.Monthly, 100.0)
	startDate := time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC)
	budgetInstance := budget.Instantiate(startDate)
	if budgetInstance.Name() != "Test Budget" {
		t.Errorf("Expected budgetInstance.Name to be 'Test Budget', got %s", budgetInstance.Name())
	}
	if budgetInstance.Amount() != 100.0 {
		t.Errorf("Expected budgetInstance.Amount to be 100.0, got %f", budgetInstance.Amount())
	}
	if budgetInstance.StartDate() != startDate {
		t.Errorf("Expected budgetInstance.StartDate to be %s, got %s", startDate, budgetInstance.StartDate())
	}
	expectedEndDate := time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC)
	if budgetInstance.EndDate() != expectedEndDate {
		t.Errorf("Expected budgetInstance.EndDate to be %s, got %s", expectedEndDate, budgetInstance.EndDate())
	}
}

func TestBudgetBiWeekly(t *testing.T) {
	budget := model.NewBudget(-1, "Test Budget", budgettype.BiWeekly, 100.0)
	startDate := time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC)
	budgetInstance := budget.Instantiate(startDate)
	if budgetInstance.Name() != "Test Budget" {
		t.Errorf("Expected budgetInstance.Name to be 'Test Budget', got %s", budgetInstance.Name())
	}
	if budgetInstance.Amount() != 100.0 {
		t.Errorf("Expected budgetInstance.Amount to be 100.0, got %f", budgetInstance.Amount())
	}
	if budgetInstance.StartDate() != startDate {
		t.Errorf("Expected budgetInstance.StartDate to be %s, got %s", startDate, budgetInstance.StartDate())
	}
	expectedEndDate := time.Date(2019, time.February, 14, 0, 0, 0, 0, time.UTC)
	if budgetInstance.EndDate() != expectedEndDate {
		t.Errorf("Expected budgetInstance.EndDate to be %s, got %s", expectedEndDate, budgetInstance.EndDate())
	}
}







