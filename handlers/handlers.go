package handlers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/moaabb/api-postgres/config"
	"github.com/moaabb/api-postgres/driver"
)

type Message struct {
	Message string `json:"message"`
}

type Handlers struct {
	L  hclog.Logger
	DB driver.DBModel
}

func NewHandlers(a *config.Application) *Handlers {
	return &Handlers{a.L, driver.NewDB(a.DBModel)}
}
