package resto

import "porto-restKelasWork/internal/model"

type UseCase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
