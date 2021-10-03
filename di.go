package wordy

import (
	"github.com/boliev/wordy/internal/controller"
	"github.com/boliev/wordy/internal/domain"
	"github.com/boliev/wordy/internal/psql"
	"github.com/boliev/wordy/internal/repository"
	"github.com/boliev/wordy/pkg/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DiCreateDB di function for database
func DiCreateDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.GetString("database_dsn")), &gorm.Config{})
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Panicf("error: %s", err.Error())
	}

	return db
}

// DiCreateConfig di function for config
func DiCreateConfig() *config.Config {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Panicf(err.Error())
	}

	return cfg
}

// DiCreateUserRepository di function for user repository
func DiCreateUserRepository(db *gorm.DB) repository.User {
	return psql.CreateUserRepository(db)
}

// DiCreateUserController di function for user controller
func DiCreateUserController(userRepository repository.User) *controller.User {
	return controller.CreateUserController(userRepository)
}

// DiCreateApp di function for app
func DiCreateApp(
	cfg *config.Config,
	db *gorm.DB,
	userRepository repository.User,
	userController *controller.User,
) *App {
	return &App{
		Cfg:            cfg,
		Db:             db,
		UserRepository: userRepository,
		UserController: userController,
	}
}
