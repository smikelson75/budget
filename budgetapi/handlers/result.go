package handlers

import "budgetlib/models"

type Result struct {
	Categories []*models.Category `json:"categories,omitempty"`
}
