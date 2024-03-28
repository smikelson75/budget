package dto

import "budgetlib/models"

type IResult[T models.Category] struct {
	Result T
}