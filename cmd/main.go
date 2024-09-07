package main

import (
	"notifier/config"
	"notifier/internal/presenter/controllers"
	"notifier/pkg/http"
	"notifier/pkg/resend"
	"notifier/pkg/validator"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		http.Module,
		validator.Module,
		// awssdk.Module,
		resend.Module,
		controllers.NotifierControllerModule,
		fx.Invoke(http.RunHttpServer),
	).Run()
}
