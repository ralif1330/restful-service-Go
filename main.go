package main

import "github.com/labstack/echo/v4"

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

	return c.JSON(201, foodMenu)
}

func main() {

	e := echo.New()
	// localhost:14045/menu/food
	e.GET("/menu/food", getFoodMenu)
	e.Logger.Fatal(e.Start(":14045"))

}
