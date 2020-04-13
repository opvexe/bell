package app

import (
	"bell/app/web"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewGinEngine,
	NewApp,
	web.NewServer,
)
