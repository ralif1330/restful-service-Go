package main

import (
	"porto-restKelasWork/internal/database"
	"porto-restKelasWork/internal/delivery/rest"
	mRepo "porto-restKelasWork/internal/repository/menu"
	rUseCase "porto-restKelasWork/internal/usecase/resto"

	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=admin dbname=resto-app-belajar  sslmode=disable"
)

func main() {

	e := echo.New()

	db := database.GetDB(dbAddress)

	menuRepo := mRepo.GetRepository(db)

	restoUseCase := rUseCase.GetUseCase(menuRepo)

	h := rest.NewHandler(restoUseCase)

	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":1404"))

}
