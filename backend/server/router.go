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
		quiz.GET("/c/:user_id", quizController.GetAllQuizFromCreatedUser)
		quiz.GET("/p/:user_id", quizController.GetAllQuizFromParticipatedUser)
		quiz.GET("/categories", quizController.GetAllQuizCategory)
		quiz.POST("/", quizController.CreateQuiz)
		quiz.PUT("/:quiz_id", quizController.UpdateQuiz)
		quiz.DELETE("/:quiz_id", quizController.DeleteQuiz)
	}

	report := api.Group("/reports")
	{
		report.GET("/", reportController.GetAllReport)
		report.GET("/:date", reportController.GetReportFromDateCreated)
		report.POST("/", reportController.CreateReport)
		report.PUT("/:report_id/accept", reportController.AcceptReport)
		report.PUT("/:report_id/complete", reportController.CompleteReport)
		report.DELETE("/:report_id", reportController.DeleteReport)
	}

	history := api.Group("/histories")
	{
		history.GET("/", historyController.GetAllHistory)
		history.GET("/:user_id", historyController.GetHistoryFromUser)
		history.POST("/", historyController.CreateHistory)
		history.DELETE("/:history_id", historyController.DeleteHistory)
		history.DELETE("/:user_id", historyController.DeleteAllHistoryFromUser)
	}

	return router
}
