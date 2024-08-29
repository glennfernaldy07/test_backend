package main

import (
	"app/bundlefx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bundlefx.CoreModules,
		bundlefx.EntityModules,
	).Run()
}
