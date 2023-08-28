package resto

import "porto-restKelasWork/internal/model"

type GetUseCase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
