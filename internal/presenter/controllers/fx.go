package controllers

import (
	"go.uber.org/fx"
)

var NotifierControllerModule = fx.Module(
	"notifier_controller",
	//USECASES
	fx.Invoke(NotifierController),
)
