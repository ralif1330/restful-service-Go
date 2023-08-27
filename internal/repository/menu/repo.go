package menu

import "porto-restKelasWork/internal/model"

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
