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
		user.PUT("/:user_id", userController.UpdateUser)
		user.DELETE("/:user_id", userController.DeleteUser)
	}

	quiz := api.Group("/quizzes")
	{
		quiz.GET("/", quizController.GetAllQuiz)
		quiz.GET("/c/:id", quizController.GetAllQuizFromCreatedUser)
		quiz.GET("/p/:id", quizController.GetAllQuizFromParticipatedUser)
		quiz.GET("/categories", quizController.GetAllQuizCategory)
		quiz.POST("/", quizController.CreateQuiz)
		quiz.PUT("/:id", quizController.UpdateQuiz)
		quiz.DELETE("/:id", quizController.DeleteQuiz)
	}

	report := api.Group("/reports")
	{
		report.GET("/", reportController.GetAllReport)
		report.GET("/:date", reportController.GetReportFromDateCreated)
		report.POST("/", reportController.CreateReport)
	}

	history := api.Group("/histories")
	{
		history.GET("/", historyController.GetAllHistory)
		history.GET("/:id", historyController.GetHistoryFromUser)
		history.POST("/", historyController.CreateHistory)
	}

	return router
}
