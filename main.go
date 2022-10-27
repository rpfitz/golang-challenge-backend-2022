package main

import (
	"backendmod/config"
	"backendmod/controller"
	"backendmod/router"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Router router.Router
}

func main() {
	app := App{}
	app.Initialize()
	app.Run(config.PORT)
}

func (a *App) Initialize() {
	a.Router = router.NewMuxRouter()

	a.InitializeRoutes()
}

func (a *App) InitializeRoutes() {
	a.Router.POST("/login", controller.Login)
	a.Router.POST("/sign-up", controller.SignUp)
	a.Router.POST("/googleauth", controller.GoogleAuthenticator)
	a.Router.GET("/authuser", controller.AuthUser)
	a.Router.POST("/edit-profile", controller.EditProfile)
	a.Router.POST("/logout", controller.Logout)
}

func (a *App) Run(port string) {
	a.Router.SERVE(port)
}
