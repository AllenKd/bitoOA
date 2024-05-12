//go:build wireinject
// +build wireinject

package main

import (
	"bitoOA/internal/config"
	"bitoOA/internal/db"
	"bitoOA/internal/logger"
	"bitoOA/internal/route/handler"
	"bitoOA/internal/user"
	"github.com/google/wire"
)

func InitializeApp() *App {
	wire.Build(NewApp, handler.New, user.New, logger.New, config.New, db.New)
	return &App{}
}
