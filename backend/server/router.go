package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lebrancconvas/FancyQuiz/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 24 * time.Hour,
	}))

	// Init Controllers.
	testController := new(controllers.TestController)
	userController := new(controllers.UserController)
	quizController := new(controllers.QuizController)
	reportController := new(controllers.ReportController)
	historyController := new(controllers.HistoryController)

	// Set API Routes.
	api := router.Group("/api")

	test := api.Group("/test")
	{
		test.GET("/ping", testController.Ping)
	}

	user := api.Group("/users")
	{
		user.GET("/", userController.GetAllUsers)
		user.POST("/", userController.CreateUser)
	}

	quiz := api.Group("/quizzes")
	{

	}

	report := api.Group("/reports")
	{

	}

	history := api.Group("/histories")
	{
		
	}

	return router
}
