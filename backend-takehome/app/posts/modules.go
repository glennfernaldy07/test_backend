package posts

import (
	"app/posts/delivery/http"
	"app/posts/service"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(http.NewServerOption),
	fx.Provide(http.NewEndpoint),
	fx.Provide(service.NewService),
	fx.Invoke(http.NewHandler),
)
