package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	"OpenAI-api/api"
)

func main() {
	// Initialize viper and read configurations
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	api.Init()

	// Create an Echo instance
	e := echo.New()

	// Set the logger to use a custom format
	e.Logger.SetLevel(log.INFO)
	e.Logger.SetOutput(os.Stdout)
	e.Logger.SetHeader("${time_rfc3339} ${level}")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("static"))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Routes
	// chat
	e.POST("/v1/chat/completions", api.HandleChat)
	e.POST("/chat/completions", api.HandleChat)

	// completions
	e.POST("/v1/completions", api.HandleCompletions)
	e.POST("/completions", api.HandleCompletions)

	// embeddings
	e.POST("/v1/embeddings", api.HandleEmbeddings)
	e.POST("/embeddings", api.HandleEmbeddings)

	// Start the server
	port := viper.GetString("port")
	if port == "" {
		port = ":8080"
	}

	e.Logger.Fatal(e.Start(port))
}
