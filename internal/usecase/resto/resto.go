package resto

import (
	"porto-restKelasWork/internal/model"
	"porto-restKelasWork/internal/repository/menu"
)

type restoUseCase struct {
	menuRepo menu.Repository
}

func (r *restoUseCase) GetMenu(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenu(menuType)
}
