package admin

import (
	"github.com/Roongkun/software-eng-ii/internal/usecase"
	"github.com/uptrace/bun"
)

type Resolver struct {
	AdminUsecase        usecase.AdminUseCase
	PhotographerUsecase usecase.PhotographerUseCase
}

func NewResolver(db *bun.DB) *Resolver {
	return &Resolver{
		AdminUsecase:        *usecase.NewAdminUseCase(db),
		PhotographerUsecase: *usecase.NewPhotographerUseCase(db),
	}
}
