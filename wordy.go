package wordy

import (
	"github.com/boliev/wordy/internal/controller"
	"github.com/boliev/wordy/internal/domain"
	"github.com/boliev/wordy/internal/psql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// App the app
type App struct {
}

// Start the app
func (app App) Start() {
	db := initDB()
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

func initDB() *gorm.DB {
	dsn := "host=localhost user=wordy password=123456 dbname=wordy port=5432 sslmode=disable TimeZone=Europe/Berlin"
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
