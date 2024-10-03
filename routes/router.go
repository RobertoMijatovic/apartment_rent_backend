package routes

import (
	"apartmentRent/handlers"
	"apartmentRent/middlewares"
	"apartmentRent/repos"
	"apartmentRent/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CorsMiddleware())

	rep := repos.NewRepos(db)
	svc := services.NewServices(rep)
	hlr := handlers.NewHandlers(svc)

	apiRouter := router.Group("/api")
	noAuthRouter := apiRouter.Group("/")
	authRouter := apiRouter.Group("/").Use(middlewares.AuthMiddleware())

	// Routes wo Auth
	// Users
	noAuthRouter.POST("register", hlr.CreateUser)
	noAuthRouter.POST("login", hlr.Login)

	// Routes w Auth
	// Users
	authRouter.GET("users", hlr.GetUsers)

	// Daily Rents
	authRouter.POST("dailyRent", hlr.CreateDailyRent)
	authRouter.GET("dailyRents", hlr.GetDailyRents)
	authRouter.GET("dailyRent/:id", hlr.GetDailyRent)
	authRouter.PUT("dailyRent", hlr.UpdateDailyRent)
	authRouter.DELETE("dailyRent/:id", hlr.DeleteDailyRent)

	// Long term rents
	authRouter.POST("longTermRent", hlr.CreateLongTermRent)
	authRouter.GET("longTermRents", hlr.GetLongTermRents)
	authRouter.GET("longTermRent/:id", hlr.GetLongTermRent)
	authRouter.PUT("longTermRent", hlr.UpdateLongTermRent)
	authRouter.DELETE("longTermRent/:id", hlr.DeleteLongTermRent)

	// Markets
	authRouter.POST("market", hlr.CreateMarket)
	authRouter.GET("markets", hlr.GetMarkets)
	authRouter.GET("market/:id", hlr.GetMarket)
	authRouter.PUT("market", hlr.UpdateMarket)
	authRouter.DELETE("market/:id", hlr.DeleteMarket)

	return router
}
