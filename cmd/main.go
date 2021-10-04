package main

import (
	"github.com/boliev/wordy"
	"go.uber.org/dig"
)

func main() {
	container := buildContainer()
	err := container.Invoke(func(app *wordy.App) {
		app.Start()
	})

	if err != nil {
		panic(err)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()
	provide(container, wordy.DiCreateConfig)
	provide(container, wordy.DiCreateDB)
	provide(container, wordy.DiCreateUserRepository)
	provide(container, wordy.DiCreateUserCreator)
	provide(container, wordy.DiCreateJwtService)
	provide(container, wordy.DiCreateUserAuthenticator)

	provide(container, wordy.DiCreateUserController)
	provide(container, wordy.DiCreateAuthController)

	provide(container, wordy.DiCreateApp)

	return container
}

func provide(container *dig.Container, constructor interface{}) {
	err := container.Provide(constructor)
	if err != nil {
		panic(err)
	}
}
