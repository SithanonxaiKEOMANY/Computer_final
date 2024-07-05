package web

import (
	"go_starter/controllers/web"
	"go_starter/routes"

	"github.com/gofiber/fiber/v2"
)

type webRoutes struct {
	controller              web.Controller
	classroomController     web.ClassRoomController
	userController          web.UserController
	userClassRoomController web.UserClassRoomController
	userBehaviorController  web.UserBehaviorController
	subjectController       web.SubjectController
}

func (w webRoutes) Install(app *fiber.App) {
	route := app.Group("web/", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	route.Post("hello", w.controller.StartController)

	//Subject
	route.Get("subjects", w.subjectController.GetSubjectController)
	route.Get("subject", w.subjectController.FilterSubjectBySubjectCodeService)
	route.Post("create-subject", w.subjectController.CreateSubjectService)
	route.Put("update-subject", w.subjectController.UpdateSubjectController)
	route.Delete("delete-subject", w.subjectController.DeleteSubjectController)

	//ClassRoom
	route.Get("get-all-classroom", w.classroomController.GetAllClassRoomControllers)
	route.Get("get-id-classroom", w.classroomController.GetClassRoomByIdControllers)
	route.Post("create-classroom", w.classroomController.CreateClassRoomControllers)
	route.Put("update-classroom", w.classroomController.UpdateClassRoomControllers)
	route.Delete("delete-classroom", w.classroomController.DeleteClassRoomControllers)

	//User
	route.Get("get-user-pagination", w.userController.GetUserTypeWithPaginationService)
	route.Get("user-info", w.userController.GetUserInfoController)

	route.Get("get-user-all", w.userController.GetAllUserController)
	route.Get("get-user-by-id/:id", w.userController.GetUserByIDController)
	route.Get("get-by-code-id", w.userController.GetUserByCodeIDController)
	route.Post("create-user", w.userController.CreateUserController)
	route.Put("update-user", w.userController.UpdateUserController)
	route.Delete("delete-user", w.userController.DeleteUserController)

	route.Get("total-subject", w.userController.CountTotalSubjectsController)
	route.Get("total-teacher", w.userController.CountTotalTeacherController)
	route.Get("total-student", w.userController.CountTotalStudentController)

	//Login
	route.Post("reset-password", w.userController.ResetPasswordController)
	route.Post("/admin/sign-in", w.userController.AdminSignInController)
	route.Post("sign-in", w.userController.SignInController)
	route.Post("sign-up", w.userController.SignUpController)

	//UserClassroom
	route.Get("get-classrooms", w.userClassRoomController.GetUserClassroomsController)
	route.Get("get-user-classroom", w.userClassRoomController.GetUserClassroomByStudentTypeController)
	route.Post("create-user-classroom", w.userClassRoomController.CreateUserClassRoomController)
	route.Delete("delete-user-classroom", w.userClassRoomController.DeleteUserClassRoomByIDController)

	//route.Get("get-user-class-room-all", w.userClassRoomController.GetAllUserClassRoomController)
	//route.Get("get-user-class-room-by-id/:id", w.userClassRoomController.GetUserClassRoomByIDController)
	//route.Get("get-by-user-id", w.userClassRoomController.GetUserClassRoomByUserIDController)
	//route.Get("get-by-class-room-id", w.userClassRoomController.GetUserClassRoomByClassRoomIDController)
	//route.Put("update-user-class-room", w.userClassRoomController.UpdateUserClassRoomController)
	//route.Delete("delete-user-class-room", w.userClassRoomController.DeleteUserClassRoomByIDController)

	//route.Get("teacher-classroom", w.userClassRoomController.GetClassroomByTeacherController)
	//route.Get("student-classroom", w.userClassRoomController.GetUserClassroomByTeacherController)

	//UserBehavior
	//route.Get("get-user-behavior-all", w.userBehaviorController.GetAllUserBehaviorController)
	//route.Get("get-user-behavior-by-id/:id", w.userBehaviorController.GetUserBehaviorByIDController)
	//route.Get("get-by-user-id", w.userBehaviorController.GetUserBehaviorByUserIDController)
	//route.Get("get-by-behavior-classroom-id", w.userBehaviorController.GetUserBehaviorByClassRoomIDController)
	////route.Post("create-user-behavior", w.userBehaviorController.CreateUserBehaviorController)
	////route.Put("update-user-behavior", w.userBehaviorController.UpdateUserBehaviorController)
	//route.Delete("delete-user-behavior", w.userBehaviorController.DeleteUserBehaviorByIDController)

	//route.Post("insert-student-behavior", w.userBehaviorController.InsertStudentBehaviorByStudentIdAndClassroomIdController)
	route.Post("check-student-behavior", w.userBehaviorController.UpdateUserBehaviorController)
	route.Post("report-user-behavior", w.userBehaviorController.ReportUserBehaviorWithPaginationController)
}

func NewWebRoutes(
	controller web.Controller,
	classroomController web.ClassRoomController,
	userController web.UserController,
	userClassRoomController web.UserClassRoomController,
	userBehaviorController web.UserBehaviorController,
	subjectController web.SubjectController,

	// controller
) routes.Routes {
	return &webRoutes{
		controller:              controller,
		classroomController:     classroomController,
		userController:          userController,
		userClassRoomController: userClassRoomController,
		userBehaviorController:  userBehaviorController,
		subjectController:       subjectController,

		//controller
	}
}
