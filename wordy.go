package wordy

import (
	"github.com/boliev/wordy/internal/controller"
	"github.com/boliev/wordy/internal/repository"
	"github.com/boliev/wordy/pkg/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// App the app
type App struct {
	Cfg            *config.Config
	Db             *gorm.DB
	UserRepository repository.User
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
			users.GET("/", app.UserController.List)
			users.GET("/:id", app.UserController.One)
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
