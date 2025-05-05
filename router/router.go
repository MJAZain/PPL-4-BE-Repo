package router

import (
	"go-gin-auth/controller"
	"go-gin-auth/internal/unit"
	"go-gin-auth/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRouter() *gin.Engine {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()

	// Group all under /api
	api := r.Group("/api")
	{
		// Auth routes
		api.POST("/users/register", controller.Register)
		api.POST("/users/login", controller.Login)
		api.POST("/users/logout", controller.Logout)

		// Protected user routes
		users := api.Group("/users")
		users.Use(middleware.AuthAdminMiddleware())
		{
			users.GET("/", controller.GetUsers)
			users.GET("/:id", controller.GetUser)
			users.PUT("/:id", controller.UpdateUser)
			users.DELETE("/:id", controller.DeleteUser)
			users.GET("/search", controller.SearchUsers)
			users.PATCH("/:id/deactivate", controller.DeactivateUser)
			users.PATCH("/:id/reactivate", controller.ReactivateUser)
			users.PUT("/:id/reset-password", controller.ResetUserPassword)
		}

		unit := unit.NewUnitHandler()
		units := api.Group("/units")
		units.Use(middleware.AuthAdminMiddleware())
		{
			units.POST("/", unit.CreateUnit)
			units.GET("/", unit.GetUnits)
			units.GET("/:id", unit.GetUnitByID)
			units.PUT("/:id", unit.UpdateUnit)
			units.DELETE("/:id", unit.DeleteUnit)
		}
	}
	// // Auth routes
	// r.POST("/register", controller.Register)
	// r.POST("/login", controller.Login)
	// r.POST("/logout", controller.Logout)

	// // Protected user routes
	// user := r.Group("/users")
	// user.Use(middleware.AuthMiddleware())
	// {
	// 	user.GET("/", controller.GetUsers)
	// 	user.GET("/:id", controller.GetUser)
	// 	user.PUT("/:id", controller.UpdateUser)
	// 	user.DELETE("/:id", controller.DeleteUser)
	// }

	return r
}
