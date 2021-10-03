package wordy

import (
	"github.com/boliev/wordy/internal/controller"
	"github.com/boliev/wordy/internal/domain"
	"github.com/boliev/wordy/internal/psql"
	"github.com/boliev/wordy/pkg/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// App the app
type App struct {
}

// Start the app
func (app App) Start() {
	cfg := createConfig()
	db := initDB(cfg.GetString("database_dsn"))
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		users := v1.Group("/users")
		{
			userController := controller.CreateUserController(psql.CreateUserRepository(db))

			users.GET("/", userController.List)
			users.GET("/:id", userController.One)
			users.POST("/", userController.Create)
		}
	}
	r.Run()
}

func initDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect to database")
	}

	var user domain.User
	err = db.AutoMigrate(&user)
	if err != nil {
		panic("cannot connect to database")
	}

	return db
}

func createConfig() *config.Config {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Panicf(err.Error())
	}

	return cfg
}
