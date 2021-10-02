package wordy

import (
	"github.com/boliev/wordy/internal/controller"
	"github.com/boliev/wordy/internal/memory"
	"github.com/gin-gonic/gin"
)

// App the app
type App struct {
}

// Start the app
func (app App) Start() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		users := v1.Group("/users")
		{
			userController := controller.CreateUserController(memory.CreateMemUserRepository())

			users.GET("/", userController.List)
			users.GET("/:id", userController.One)
			users.POST("/", userController.Create)
		}
	}
	r.Run()
}
