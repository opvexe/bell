//+build wireinject

package cmd

import (
	"bell/app"
	"bell/config"
	"bell/controller"
	"bell/library/database"
	"bell/library/logger"
	"bell/middleware"
	"bell/repository"
	"bell/router"
	"bell/service"
	"github.com/google/wire"
)

func BuildApp(path string) (*app.App,error) {
	panic(wire.Build(
		config.WireSet,
		logger.New,
		database.WireSet,
		repository.WireSet,
		service.WireSet,
		controller.WireSet,
		middleware.WireSet,
		router.WireSet,
		app.WireSet,
	))
}