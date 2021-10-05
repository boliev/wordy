package wordy

import (
	"github.com/boliev/wordy/internal/controller"
	"github.com/boliev/wordy/internal/middleware"
	"github.com/gin-gonic/gin"
)

// App the app
type App struct {
	AuthHandler    *middleware.AuthHandler
	UserController *controller.User
	AuthController *controller.Auth
}

// Start the app
func (app App) Start() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", app.AuthHandler.Handle, app.UserController.List)
			users.GET("/:id", app.AuthHandler.Handle, app.UserController.One)
			users.POST("/", app.UserController.Create)
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/", app.AuthController.Auth)
		}
	}
	err := r.Run()
	if err != nil {
		panic(err.Error())
	}
}
