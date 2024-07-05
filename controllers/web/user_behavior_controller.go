package web

import (
	"go_starter/controllers"
	"go_starter/logs"
	"go_starter/requests"
	"go_starter/services"
	"go_starter/trails"
	"go_starter/validation"

	"github.com/gofiber/fiber/v2"
)

type UserBehaviorController interface {
	ReportUserBehaviorWithPaginationController(ctx *fiber.Ctx) error
	UpdateUserBehaviorController(ctx *fiber.Ctx) error

	//InsertStudentBehaviorByStudentIdAndClassroomIdController(ctx *fiber.Ctx) error

	//GetAllUserBehaviorController(ctx *fiber.Ctx) error
	//GetUserBehaviorByIDController(ctx *fiber.Ctx) error
	//GetUserBehaviorByUserIDController(ctx *fiber.Ctx) error
	//GetUserBehaviorByClassRoomIDController(ctx *fiber.Ctx) error
	////CreateUserBehaviorController(ctx *fiber.Ctx) error
	//DeleteUserBehaviorByIDController(ctx *fiber.Ctx) error
}

type userBehaviorController struct {
	serviceUserBehavior services.UserBehaviorService
}

func (u *userBehaviorController) ReportUserBehaviorWithPaginationController(ctx *fiber.Ctx) error {
	request := new(requests.UserBehaviorWithPaginationRequest)
	if err := ctx.BodyParser(request); err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	paginateRequest := trails.PaginateRequest{
		Item:    request.PerPage,
		Page:    request.CurrentPage,
		Sorting: request.Sorting,
	}
	response, err := u.serviceUserBehavior.ReportUserBehaviorWithPaginationService(requests.UserBehaviorWithPaginationRequest{
		PerPage:     request.PerPage,
		CurrentPage: request.CurrentPage,
		Sorting:     request.Sorting,
		UserType:    request.UserType,
		ClassroomID: request.ClassroomID,
		SubjectID:   request.SubjectID,
	}, paginateRequest,
	)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	//return controllers.NewSuccessResponse(ctx, response)
	return ctx.JSON(response)
}

func (u *userBehaviorController) UpdateUserBehaviorController(ctx *fiber.Ctx) error {

	request := new(requests.UserBehaviorRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := u.serviceUserBehavior.UpdateUserBehaviorService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessMsg(ctx, response.Message)
}

//func (u *userBehaviorController) InsertStudentBehaviorByStudentIdAndClassroomIdController(ctx *fiber.Ctx) error {
//	request := new(requests.StudentBehaviorRequest)
//	if err := ctx.BodyParser(request); err != nil {
//		return controllers.NewErrorResponses(ctx, err)
//	}
//	errValidate := validation.Validate(request)
//	if errValidate != nil {
//		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
//	}
//	response, err := u.serviceUserBehavior.InsertStudentBehaviorByStudentIdAndClassroomIdService(*request)
//	if err != nil {
//		return controllers.NewErrorResponses(ctx, err)
//	}
//	return controllers.NewSuccessMsg(ctx, response)
//}

//func (u *userBehaviorController) CreateUserBehaviorController(ctx *fiber.Ctx) error {
//
//	request := new(requests.UserBehaviorRequest)
//	if err := ctx.BodyParser(request); err != nil {
//		return controllers.NewErrorResponses(ctx, err)
//	}
//	errValidate := validation.Validate(request)
//	if errValidate != nil {
//		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
//	}
//	response, err := u.serviceUserBehavior.CreateUserBehaviorService(*request)
//	if err != nil {
//		return controllers.NewErrorResponses(ctx, err)
//	}
//	return controllers.NewSuccessMsg(ctx, response.Message)
//}

//func (u *userBehaviorController) DeleteUserBehaviorByIDController(ctx *fiber.Ctx) error {
//
//	request := new(requests.UserBehaviorClassRoomByUserIDRequest)
//	if err := ctx.BodyParser(request); err != nil {
//		return controllers.NewErrorResponses(ctx, err)
//	}
//	errValidate := validation.Validate(request)
//	if errValidate != nil {
//		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
//	}
//	response, err := u.serviceUserBehavior.DeleteUserBehaviorService(*request)
//	if err != nil {
//		return controllers.NewErrorResponses(ctx, err)
//	}
//	return controllers.NewSuccessMsg(ctx, response.Message)
//}
//
//func (u *userBehaviorController) GetAllUserBehaviorController(ctx *fiber.Ctx) error {
//
//	customers, err := u.serviceUserBehavior.GetAllUserBehaviorServices()
//	if err != nil {
//		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"success": false,
//			"message": "Failed to retrieve customer data",
//			"error":   err.Error(),
//		})
//	}
//
//	//return http response
//	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
//		"success": true,
//		"data":    customers,
//	})
//}

//func (u *userBehaviorController) GetUserBehaviorByClassRoomIDController(ctx *fiber.Ctx) error {
//
//	req := new(requests.UserBehaviorClassRoomByIDRequest)
//	if err := ctx.BodyParser(req); err != nil {
//		logs.Error(err)
//		return controllers.NewErrorResponses(ctx, err)
//	}
//	errValidate := validation.Validate(req)
//	if errValidate != nil {
//		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
//	}
//	response, err := u.serviceUserBehavior.GetByUserBehaviorIdService(req.ClassRoomID)
//	if err != nil {
//		return controllers.NewErrorResponses(ctx, err)
//	}
//	return controllers.NewSuccessResponse(ctx, response)
//}
//
//func (u *userBehaviorController) GetUserBehaviorByIDController(ctx *fiber.Ctx) error {
//
//	id, err := ctx.ParamsInt("id", 1)
//	if err != nil {
//		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"success": false,
//			"message": "Failed to retrieve customer data",
//			"error":   err.Error(),
//		})
//	}
//	response, err := u.serviceUserBehavior.GetByIdUserBehaviorService(uint(id))
//	if err != nil {
//		return controllers.NewErrorResponses(ctx, err)
//	}
//	return controllers.NewSuccessResponse(ctx, response)
//}
//
//func (u *userBehaviorController) GetUserBehaviorByUserIDController(ctx *fiber.Ctx) error {
//
//	req := new(requests.UserBehaviorClassRoomByUserIDRequest)
//	if err := ctx.BodyParser(req); err != nil {
//		logs.Error(err)
//		return controllers.NewErrorResponses(ctx, err)
//	}
//	errValidate := validation.Validate(req)
//	if errValidate != nil {
//		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
//	}
//	response, err := u.serviceUserBehavior.GetByUserIDService(req.UserID)
//	if err != nil {
//		return controllers.NewErrorResponses(ctx, err)
//	}
//	return controllers.NewSuccessResponse(ctx, response)
//}

func NewUserBehaviorController(serviceUserBehavior services.UserBehaviorService) UserBehaviorController {

	return &userBehaviorController{serviceUserBehavior: serviceUserBehavior}
}
