package dto

import "github.com/Eri-Vadi/go_dining_hall/internal/domain/entity"

type Menu struct {
	ItemsCount int           `json:"items_count"`
	Items      []entity.Food `json:"items"`
}
