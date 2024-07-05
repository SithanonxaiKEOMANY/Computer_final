package api

import (
	"go_starter/controllers/api"
	"go_starter/routes"

	"github.com/gofiber/fiber/v2"
)

type apiRoutes struct {
	controllerApi api.ControllerApi
}

func (a apiRoutes) Install(app *fiber.App) {
	route := app.Group("api/", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	route.Post("hello", a.controllerApi.StartController)

}

func NewApiRoutes(
	controllerApi api.ControllerApi,

	// controller
) routes.Routes {
	return &apiRoutes{
		controllerApi: controllerApi,
		//controller
	}
}
