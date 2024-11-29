package main

import (
	config "authentication-service/configs"
	"authentication-service/internal/handler"
	"authentication-service/internal/repository"
	"authentication-service/internal/server"
	"authentication-service/internal/service"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

const configPath = "./configs"

func main() {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("wrong config variables")
	}

	db, err := newPostgresDB(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("err initializing DB")
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	srv := server.NewServer(cfg, handlers.InitRoutes())

	go func() {
		if err := srv.Run(); err != http.ErrServerClosed {
			log.Error().Err(err).Msg("error occurred while running http server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to stop server")
	}
}

func newPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	return repository.NewPostgresDB(repository.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})
}
