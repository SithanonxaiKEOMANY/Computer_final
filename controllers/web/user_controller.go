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

type UserController interface {
	GetUserInfoController(ctx *fiber.Ctx) error

	GetUserTypeWithPaginationService(ctx *fiber.Ctx) error
	CountTotalSubjectsController(ctx *fiber.Ctx) error
	CountTotalTeacherController(ctx *fiber.Ctx) error
	CountTotalStudentController(ctx *fiber.Ctx) error

	//admin
	ResetPasswordController(ctx *fiber.Ctx) error
	AdminSignInController(ctx *fiber.Ctx) error
	SignInController(ctx *fiber.Ctx) error
	SignUpController(ctx *fiber.Ctx) error

	GetAllUserController(ctx *fiber.Ctx) error
	GetUserByIDController(ctx *fiber.Ctx) error
	GetUserByCodeIDController(ctx *fiber.Ctx) error
	CreateUserController(ctx *fiber.Ctx) error
	UpdateUserController(ctx *fiber.Ctx) error
	DeleteUserController(ctx *fiber.Ctx) error
}

type userController struct {
	serviceUser services.UserService
}

func (u *userController) GetUserInfoController(ctx *fiber.Ctx) error {
	req := new(requests.UserIdRequest)
	if err := ctx.BodyParser(req); err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(req)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := u.serviceUser.GetUserInfoService(*req)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (u *userController) ResetPasswordController(ctx *fiber.Ctx) error {
	request := new(requests.ResetPasswordRequest)
	if err := ctx.BodyParser(request); err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := u.serviceUser.ResetUserPasswordService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (u *userController) GetUserTypeWithPaginationService(ctx *fiber.Ctx) error {
	request := new(requests.UserWithPaginationRequest)
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
	response, err := u.serviceUser.GetUserTypeWithPaginationService(requests.UserWithPaginationRequest{
		PerPage:     request.PerPage,
		CurrentPage: request.CurrentPage,
		Sorting:     request.Sorting,
		UserType:    request.UserType,
	}, paginateRequest,
	)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	//return controllers.NewSuccessResponse(ctx, response)
	return ctx.JSON(response)
}

func (u *userController) CountTotalSubjectsController(ctx *fiber.Ctx) error {
	response, err := u.serviceUser.CountTotalSubjectsService()
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (u *userController) CountTotalTeacherController(ctx *fiber.Ctx) error {
	response, err := u.serviceUser.CountTotalTeachersService()
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (u *userController) CountTotalStudentController(ctx *fiber.Ctx) error {
	response, err := u.serviceUser.CountTotalStudentService()
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (u *userController) CreateUserController(ctx *fiber.Ctx) error {

	request := new(requests.UserRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := u.serviceUser.CreateUserService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessMsg(ctx, response.Message)
}

func (u *userController) DeleteUserController(ctx *fiber.Ctx) error {

	request := new(requests.UserCodeIdRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := u.serviceUser.DeleteUserService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessMsg(ctx, response.Message)
}

func (u *userController) GetAllUserController(ctx *fiber.Ctx) error {

	customers, err := u.serviceUser.GetAllUserService()
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

func (u *userController) GetUserByCodeIDController(ctx *fiber.Ctx) error {

	req := new(requests.UserCodeIdRequest)
	if err := ctx.BodyParser(req); err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(req)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := u.serviceUser.GetByCodeIdUserService(*req)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (u *userController) GetUserByIDController(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id", 1)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve customer data",
			"error":   err.Error(),
		})
	}
	response, err := u.serviceUser.GetByIdUserService(uint(id))
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (u *userController) AdminSignInController(ctx *fiber.Ctx) error {

	req := new(requests.SignInUserRequest)
	if err := ctx.BodyParser(req); err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(req)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := u.serviceUser.AdminSignInUserService(*req)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (u *userController) SignInController(ctx *fiber.Ctx) error {

	req := new(requests.SignInUserRequest)
	if err := ctx.BodyParser(req); err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(req)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := u.serviceUser.SignInUserService(*req)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (u *userController) SignUpController(ctx *fiber.Ctx) error {

	req := new(requests.SigUpUserRequest)
	if err := ctx.BodyParser(req); err != nil {
		logs.Error(err)
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(req)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := u.serviceUser.SignUpUserService(*req)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (u *userController) UpdateUserController(ctx *fiber.Ctx) error {

	request := new(requests.UserRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := u.serviceUser.UpdateUserService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessMsg(ctx, response.Message)
}

func NewUserController(serviceUser services.UserService) UserController {

	return &userController{serviceUser: serviceUser}
}
