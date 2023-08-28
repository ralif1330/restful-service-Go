package rest

import "porto-restKelasWork/internal/usecase/resto"

type handler struct {
	restoUseCase resto.GetUseCase
}

func NewHandler(restoUseCase resto.GetUseCase) *handler {
	return &handler{
		restoUseCase: restoUseCase,
	}
}
