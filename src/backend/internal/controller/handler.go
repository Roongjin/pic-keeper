package controller

import (
	"github.com/Roongkun/software-eng-ii/internal/controller/user"
	"github.com/uptrace/bun"
)

type Handler struct {
	User user.Resolver
}

func NewHandler(db *bun.DB) *Handler {
	return &Handler{
		User: *user.NewResolver(db),
	}
}