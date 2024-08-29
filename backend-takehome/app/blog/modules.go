package blog

import (
	"app/blog/repository/db"
	"app/blog/service"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(service.NewService),
	fx.Provide(db.NewRepository),
)
