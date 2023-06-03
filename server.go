package main

import (
	"os"
	"sentiment/config"
	"sentiment/controllers"
	"sentiment/middleware"
	"sentiment/repository"
	"sentiment/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                     *gorm.DB                           = config.SetupDatabaseConnection()
	userRepository         repository.UserRepository          = repository.NewUserRepository(db)
	studyProgramRepository repository.StudyProgramRepository  = repository.NewStudyProgramRepository(db)
	jwtService             service.JWTService                 = service.NewJWTService()
	authService            service.AuthService                = service.NewAuthService(userRepository)
	userService            service.UserService                = service.NewUserService(userRepository)
	studyProgramService    service.StudyProgramService        = service.NewStudyProgramService(studyProgramRepository)
	authController         controllers.AuthController         = controllers.NewAuthController(authService, jwtService)
	userController         controllers.UserController         = controllers.NewUserController(userService, jwtService)
	studyProgramController controllers.StudyProgramController = controllers.NewStudyProgramController(studyProgramService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	authRoute := router.Group("auth")
	{
		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register", authController.Register)
	}

	userRoutes := router.Group("users", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	studyProgramRoutes := router.Group("study-program", middleware.AuthorizeJWT(jwtService))
	{
		studyProgramRoutes.POST("/", studyProgramController.Create)
		studyProgramRoutes.PUT("/", studyProgramController.Update)
		studyProgramRoutes.DELETE("/:id", studyProgramController.Delete)
		studyProgramRoutes.GET("/", studyProgramController.All)
		studyProgramRoutes.GET("/:id", studyProgramController.FindByID)
		studyProgramRoutes.GET("/code/:code", studyProgramController.FindByCode)
	}

	router.Run(":" + port)
}
