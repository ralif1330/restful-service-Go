package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(c echo.Context) error {

	foodMenu := []MenuItem{
		{
			Name:      "Tempe",
			OrderCode: "tempe",
			Price:     5000,
		},
		{
			Name:      "Terong",
			OrderCode: "terong",
			Price:     7000,
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data" : foodMenu,
	})
}

func getDrinksMenu(c echo.Context) error {

	drinkMenu := []MenuItem{
		{
			Name:      "Milo",
			OrderCode: "milo",
			Price:     3000,
		},
		{
			Name:      "Susu",
			OrderCode: "susu",
			Price:     2000,
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data" : drinkMenu,
	})
}

func main() {

	e := echo.New()
	// localhost:14045/menu/food
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drinks", getDrinksMenu)

	e.Logger.Fatal(e.Start(":14045"))

}
