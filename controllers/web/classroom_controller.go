package web

import (
	"fmt"
	"go_starter/controllers"
	"go_starter/requests"
	"go_starter/services"
	"go_starter/validation"

	"github.com/gofiber/fiber/v2"
)

type ClassRoomController interface {
	GetAllClassRoomControllers(ctx *fiber.Ctx) error
	GetClassRoomByIdControllers(ctx *fiber.Ctx) error
	CreateClassRoomControllers(ctx *fiber.Ctx) error
	UpdateClassRoomControllers(ctx *fiber.Ctx) error
	DeleteClassRoomControllers(ctx *fiber.Ctx) error
}

type classroomController struct {
	serviceClassRoom services.ClassRoomService
}

func (c *classroomController) CreateClassRoomControllers(ctx *fiber.Ctx) error {
	request := new(requests.CreateClassRoomRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	fmt.Println(request)
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := c.serviceClassRoom.CreateClassRoomService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}

	return controllers.NewSuccessMessage(ctx, response.Message)
}


func (c *classroomController) GetAllClassRoomControllers(ctx *fiber.Ctx) error {
	response, err := c.serviceClassRoom.GetAllClassRoomServices()
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (c *classroomController) GetClassRoomByIdControllers(ctx *fiber.Ctx) error {
	request := new(requests.ClassRoomIDRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := c.serviceClassRoom.GetClassRoomByIdService(request.Id)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}


func (c *classroomController) DeleteClassRoomControllers(ctx *fiber.Ctx) error {
	request := new(requests.DeleteClassRoomRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := c.serviceClassRoom.DeleteClassRoomService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessMsg(ctx, response.Message)
}
func (c *classroomController) UpdateClassRoomControllers(ctx *fiber.Ctx) error {
	request := new(requests.UpdateClassRoomRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := c.serviceClassRoom.UpdateClassRoomService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessMsg(ctx, response.Message)
}

func NewRoomController(serviceClassRoom services.ClassRoomService) ClassRoomController {
	return &classroomController{serviceClassRoom: serviceClassRoom}
}
