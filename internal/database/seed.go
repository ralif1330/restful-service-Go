package database

import (
	"fmt"
	"porto-restKelasWork/internal/model"
	"porto-restKelasWork/internal/model/constant"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {

	foodMenu := []model.MenuItem{
		{
			Name:      "Tempe",
			OrderCode: "tempe",
			Price:     5000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Terong",
			OrderCode: "terong",
			Price:     7000,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name:      "Milo",
			OrderCode: "milo",
			Price:     3000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Susu",
			OrderCode: "susu",
			Price:     2000,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}

}
