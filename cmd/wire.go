//+build wireinject

package cmd

import (
	conf "bell/config"
	ctrl "bell/controller"
	"github.com/google/wire"
	app "bell/app"
	api "bell/router"
	mid "bell/middleware"
	svc "bell/service"
	repos "bell/repository"
)

func BuildApp(path string) (*app.App,error) {
	panic(wire.Build(
		conf.WireSet,
		repos.WireSet,
		ctrl.WireSet,
		svc.WireSet,
		mid.WireSet,
		api.WireSet,
		app.WireSet,
	))
}