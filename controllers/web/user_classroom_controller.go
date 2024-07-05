package web

import (
	"go_starter/controllers"
	"go_starter/logs"
	"go_starter/requests"
	"go_starter/services"
	"go_starter/validation"

	"github.com/gofiber/fiber/v2"
)

type UserClassRoomController interface {
	GetUserClassroomsController(ctx *fiber.Ctx) error
	GetUserClassroomByStudentTypeController(ctx *fiber.Ctx) error
	CreateUserClassRoomController(ctx *fiber.Ctx) error

	GetClassroomByTeacherController(ctx *fiber.Ctx) error
	GetAllUserClassRoomController(ctx *fiber.Ctx) error
	GetUserClassRoomByIDController(ctx *fiber.Ctx) error
	GetUserClassRoomByUserIDController(ctx *fiber.Ctx) error
	GetUserClassRoomByClassRoomIDController(ctx *fiber.Ctx) error
	UpdateUserClassRoomController(ctx *fiber.Ctx) error
	DeleteUserClassRoomByIDController(ctx *fiber.Ctx) error
}

type userClassRoomController struct {
	serviceUserClassRoom services.UserClassRoomService
}

func (uc *userClassRoomController) GetUserClassroomsController(ctx *fiber.Ctx) error {
	request := new(requests.GetUserClassroomRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := uc.serviceUserClassRoom.GetUserClassroomsService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (uc *userClassRoomController) GetUserClassroomByStudentTypeController(ctx *fiber.Ctx) error {
	request := new(requests.UserClassroomRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := uc.serviceUserClassRoom.GetUserClassroomByStudentTypeService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (uc *userClassRoomController) GetClassroomByTeacherController(ctx *fiber.Ctx) error {
	request := new(requests.TeacherIdRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := uc.serviceUserClassRoom.GetClassroomByTeacherService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (uc *userClassRoomController) CreateUserClassRoomController(ctx *fiber.Ctx) error {
	var request requests.CreateUserClassroomANDUserBehaviorRoomRequest

	if err := ctx.BodyParser(&request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}

	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}

	response, err := uc.serviceUserClassRoom.CreateUserClassRoomAndUserBehaviorService(request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}

	return controllers.NewSuccessMsg(ctx, response.Message)
}

//func (uc *userClassRoomController) CreateUserClassRoomController(ctx *fiber.Ctx) error {
//
//	//request := new(requests.UserClassRoomRequest)
//	var request []requests.UserClassRoomRequest
//	if err := ctx.BodyParser(request); err != nil {
//		return controllers.NewErrorResponses(ctx, err)
//	}
//	errValidate := validation.Validate(request)
//	if errValidate != nil {
//		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
//	}
//	response, err := uc.serviceUserClassRoom.CreateUserClassRoomService(request)
//	if err != nil {
//		return controllers.NewErrorResponses(ctx, err)
//	}
//	return controllers.NewSuccessMsg(ctx, response.Message)
//}

func (uc *userClassRoomController) DeleteUserClassRoomByIDController(ctx *fiber.Ctx) error {

	request := new(requests.DeleteUserClassRoomByUserIDRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := uc.serviceUserClassRoom.DeleteUserClassRoomService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessMsg(ctx, response.Message)
}

func (uc *userClassRoomController) GetAllUserClassRoomController(ctx *fiber.Ctx) error {

	customers, err := uc.serviceUserClassRoom.GetAllUserClassRoomServices()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve customer data",
			"error":   err.Error(),
		})
	}

	//return http response
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    customers,
	})
}

func (uc *userClassRoomController) GetUserClassRoomByClassRoomIDController(ctx *fiber.Ctx) error {

	req := new(requests.UserClassRoomByIDRequest)
	if err := ctx.BodyParser(req); err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(req)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := uc.serviceUserClassRoom.GetByUserClassRoomIdService(req.ClassRoomID)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (uc *userClassRoomController) GetUserClassRoomByIDController(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id", 1)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve customer data",
			"error":   err.Error(),
		})
	}
	response, err := uc.serviceUserClassRoom.GetByIdUserClassRoomService(uint(id))
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (uc *userClassRoomController) GetUserClassRoomByUserIDController(ctx *fiber.Ctx) error {

	req := new(requests.UserClassRoomByUserIDRequest)
	if err := ctx.BodyParser(req); err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(req)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := uc.serviceUserClassRoom.GetByUserIDService(req.UserID)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (uc *userClassRoomController) UpdateUserClassRoomController(ctx *fiber.Ctx) error {

	request := new(requests.UpdateUserClassRoomRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := uc.serviceUserClassRoom.UpdateUserClassRoomService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessMsg(ctx, response.Message)
}

func NewUserClassRoomController(serviceUserClassRoom services.UserClassRoomService) UserClassRoomController {

	return &userClassRoomController{serviceUserClassRoom: serviceUserClassRoom}
}
