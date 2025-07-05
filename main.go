package main

import (
	"bwastartup/config"
	"bwastartup/database"
	"bwastartup/handler"
	"bwastartup/user"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	newConfig := config.NewConfig()

	db, err := database.ConnectionPostgres(newConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to DB")
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()
}
