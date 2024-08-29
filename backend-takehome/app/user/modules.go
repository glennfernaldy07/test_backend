package user

import (
	"app/user/delivery/http"
	"app/user/repository/cache"
	"app/user/repository/db"
	"app/user/service"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(http.NewServerOption),
	fx.Provide(http.NewEndpoint),
	fx.Provide(service.NewService),
	fx.Provide(db.NewRepository),
	fx.Provide(cache.NewRepository),
	fx.Invoke(http.NewHandler),
)
