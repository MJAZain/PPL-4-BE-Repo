package router

import (
	"go-gin-auth/controller"
	"go-gin-auth/middleware"
	"go-gin-auth/repository"
	"go-gin-auth/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	repo := repository.NewTransaksiRepository()
	svc := service.NewTransaksiService(repo)
	//ctrl := controller.NewTransaksiController(svc)
	api := r.Group("/api")
	{
		api.POST("/users/login", controller.Login)
		users := api.Group("/users")
		users.Use(middleware.AuthAdminMiddleware())
		{
			users.POST("/register", controller.Register)
			users.POST("/logout", controller.Logout)
			users.GET("/", controller.GetUsers)
			users.GET("/:id", controller.GetUser)
			users.PUT("/:id", controller.UpdateUser)
			users.DELETE("/:id", controller.DeleteUser)
			users.GET("/search", controller.SearchUsers)
			users.PATCH("/:id/deactivate", controller.DeactivateUser)
			users.PATCH("/:id/reactivate", controller.ReactivateUser)
			users.PUT("/:id/reset-password", controller.ResetUserPassword)
			users.GET("/export/csv", controller.ExportUsersCSV)

		}
		transaksi := api.Group("/transaksi")
		transaksi.Use(middleware.AuthAdminMiddleware()).DELETE("/:id", controller.NewTransaksiController(svc).DeleteTransaksi)
		transaksi.Use(middleware.AuthMiddleware())
		{
			transaksi.POST("/", controller.NewTransaksiController(svc).CreateTransaksi)
			transaksi.GET("/", controller.NewTransaksiController(svc).GetAllTransaksi)
		}
	}
	return r
}
