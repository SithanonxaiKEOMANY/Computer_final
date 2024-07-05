package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"

	"go_starter/config"
	"go_starter/controllers/api"
	"go_starter/controllers/web"
	"go_starter/database"
	"go_starter/logs"
	"go_starter/repositories"
	apiRoute "go_starter/routes/api"
	webRoute "go_starter/routes/web"
	"go_starter/services"
)

func main() {
	// connect database
	// postgres
	postgresConnection, err := database.PostgresConnection()
	if err != nil {
		logs.Error(err)
		return
	}

	////call api client interface
	//httpClient := http.Client{}
	//newHttpClientTrail := trails.NewHttpClientTrail(httpClient)
	//partners.NewPartner(newHttpClientTrail)

	//basic structure
	newRepository := repositories.NewRepository(postgresConnection)
	newService := services.NewService(newRepository)
	newControllerApi := api.NewControllerApi(newService)

	//Subject
	subjectRepository := repositories.NewSubjectRepository(postgresConnection)
	subjectService := services.NewSubjectService(subjectRepository)
	subjectController := web.NewSubjectController(subjectService)

	//Classroom
	classroomRepository := repositories.NewRoomRepository(postgresConnection)
	classroomService := services.NewRoomServices(classroomRepository)
	classroomController := web.NewRoomController(classroomService)

	//User
	UserRepository := repositories.NewUserRepository(postgresConnection)
	userService := services.NewUserService(UserRepository)
	userController := web.NewUserController(userService)

	//UserBehavior
	userBehaviorRepository := repositories.NewUserBehaviorRepository(postgresConnection)
	userBehaviorService := services.NewUserBehaviorService(userBehaviorRepository)
	userBehaviorController := web.NewUserBehaviorController(userBehaviorService)

	//UserClassroom
	userClassRoomRepository := repositories.NewUserClassRoom(postgresConnection)
	userClassRoomService := services.NewUserClassRoomService(userClassRoomRepository, userBehaviorService)
	userClassRoomController := web.NewUserClassRoomController(userClassRoomService)

	//connect route
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   16 * 1024 * 1024,
	})
	app.Use(logger.New())
	app.Use(cors.New())

	//Web routes
	newController := web.NewController(newService)
	newWebRoute := webRoute.NewWebRoutes(
		newController,
		classroomController,
		userController,
		userClassRoomController,
		userBehaviorController,
		subjectController,
		//new web controller
	)
	newWebRoute.Install(app)

	//Api routes
	newApiRoute := apiRoute.NewApiRoutes(
		newControllerApi, // Using the previously declared instance
	)
	newApiRoute.Install(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env("app.port"))))
}
