package main

import (
    "github.com/TgrRys/finalProject-BTPN/database"
    "github.com/TgrRys/finalProject-BTPN/controllers"
    "github.com/TgrRys/finalProject-BTPN/services"
    "github.com/TgrRys/finalProject-BTPN/helpers"
    "github.com/TgrRys/finalProject-BTPN/repositories"
    "github.com/TgrRys/finalProject-BTPN/routers"
    "fmt"
)

func main() {
	db := database.DBConnect()

	authService := helpers.NewAuthService()

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepository)

	controller := &routers.Controllers{
		UserController:  controllers.NewUserController(userService, authService),
		PhotoController: controllers.NewPhotoController(photoService, authService),
	}

	middleware := &routers.AuthMiddleware{
		AuthService: authService,
		UserService: userService,
	}

	r := routers.NewRouter(controller, middleware)

	err := r.Run(":5000")
	if err != nil {
		fmt.Println("Error on the route run")
	}
}