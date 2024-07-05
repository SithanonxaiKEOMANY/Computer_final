package web

import (
	"github.com/gofiber/fiber/v2"
	"go_starter/controllers"
	"go_starter/services"
)

type Controller interface {
	StartController(ctx *fiber.Ctx) error
}

type controller struct {
	service services.Service
}

func (c controller) StartController(ctx *fiber.Ctx) error {
	//TODO implement me
	return controllers.NewSuccessMsg(ctx, "Hello Golang web api")
}

func NewController(
	service services.Service,
	// services
) Controller {
	return &controller{
		service: service,
		//services
	}
}
