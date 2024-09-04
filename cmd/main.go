package main

import (
	"notifier/config"
	"notifier/internal/presenter/controllers"
	awssdk "notifier/pkg/aws"
	"notifier/pkg/http"
	"notifier/pkg/validator"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		http.Module,
		validator.Module,
		awssdk.Module,
		controllers.NotifierControllerModule,
		fx.Invoke(http.RunHttpServer),
	).Run()
}
