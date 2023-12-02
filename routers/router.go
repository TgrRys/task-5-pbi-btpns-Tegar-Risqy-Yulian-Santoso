package routers

import (
    "github.com/TgrRys/finalProject-BTPN/controllers"
    "github.com/TgrRys/finalProject-BTPN/middlewares"
    "github.com/TgrRys/finalProject-BTPN/helpers"
    "github.com/TgrRys/finalProject-BTPN/services"
    "github.com/gin-gonic/gin"
)

type Controllers struct {
	UserController  controllers.UserController
	PhotoController controllers.PhotoController
}

type AuthMiddleware struct {
	AuthService helpers.JwtService
	UserService services.UserService
}

func NewRouter(c *Controllers, a *AuthMiddleware) *gin.Engine {
	router := gin.Default()
	api := router.Group("api/v1")

	userRoute := api.Group("/users")
	{
		userRoute.POST("/register", c.UserController.Register)
		userRoute.POST("/login", c.UserController.Login)
		userRoute.PUT("/:userId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.UserController.Update)
		userRoute.DELETE("/:userId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.UserController.Delete)
	}

	photoRoute := api.Group("/photos")
	{
		photoRoute.POST("/", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.Create)
		photoRoute.GET("/", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.GetPhoto)
		photoRoute.GET("/:photoId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.GetPhotoByID)
		photoRoute.PUT("/:photoId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.Edit)
		photoRoute.DELETE("/:photoId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.Delete)
	}

	return router
}