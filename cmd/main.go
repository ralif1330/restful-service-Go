package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

const (
	MenuTypeFood  = "food"
	MenuTypeDrink = "drink"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=admin dbname=resto-app-belajar sslmode=disable"
)

type MenuType string

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}

func main() {

	seedDB()

	e := echo.New()
	// localhost:14045/menu/food
	e.GET("/menu", getMenu)

	e.Logger.Fatal(e.Start(":1404"))

}



func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}
	var menuData []MenuItem

	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})

}