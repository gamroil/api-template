package api

import (
	"api-template/api/observability"
	"fmt"
	"os"

	"api-template/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type App struct {
	logger  *zap.Logger
	config  *config.Configuration
	metrics *observability.Client
}

func New(configPath string) *App {
	if configPath == "" {
		configPath = "config.yaml"
	}

	logger := zap.Must(zap.NewProduction())
	defer func() {
		_ = logger.Sync()
	}()

	return &App{
		logger:  logger,
		config:  config.Configure(configPath),
		metrics: observability.New(logger),
	}
}

// Start runs the API. It is triggered by the serve command.
func (a *App) Start() {
	a.run()
}

func (a *App) run() {
	router := gin.Default()
	a.createRoutes(router)

	port := fmt.Sprintf(":%d", a.config.Port)
	if err := router.Run(port); err != nil {
		a.logger.Error("starting the API: ", zap.Error(err))
		os.Exit(1)
	}
}
