package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ayserragm/backend-project/internal/config"
	"github.com/ayserragm/backend-project/internal/db"
	"github.com/ayserragm/backend-project/internal/logger"
	"github.com/ayserragm/backend-project/internal/routes"
)

func main() {
	cfg := config.LoadConfig()

	logger.InitLogger()
	logger.Log.Info().Msg("🚀 Application starting...")

	db.ConnectDatabase(cfg)

	router := gin.Default()
	routes.RegisterRoutes(router, cfg)

	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}

	go func() {
		logger.Log.Info().Msgf("%s running on port %s", cfg.AppName, cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal().Err(err).Msg("Server error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Log.Info().Msg("🛑 Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal().Err(err).Msg("Forced shutdown")
	}

	logger.Log.Info().Msg("✅ Shutdown complete")
}
