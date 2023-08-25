package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func seedDB() {

	foodMenu := []MenuItem{
		{
			Name:      "Tempe",
			OrderCode: "tempe",
			Price:     5000,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Terong",
			OrderCode: "terong",
			Price:     7000,
			Type:      MenuTypeFood,
		},
	}

	drinkMenu := []MenuItem{
		{
			Name:      "Milo",
			OrderCode: "milo",
			Price:     3000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Susu",
			OrderCode: "susu",
			Price:     2000,
			Type:      MenuTypeDrink,
		},
	}

	fmt.Println(drinkMenu, foodMenu)

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&MenuItem{})

	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}

}