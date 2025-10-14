package config

import (
	"log"
	"log/slog"

	"github.com/kennedymwaniki/events-api/internal/infrastructure"
	"github.com/kennedymwaniki/events-api/internal/infrastructure/pg"

	usecase "github.com/kennedymwaniki/events-api/internal/usecase/event"
)

type Config struct {
	logger *slog.Logger
	Event  *usecase.EventUseCase
}

var Appconfig *Config

func InitializeDependencies() *Config {
	db, err := pg.NewPool()
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %v", err)
	}
	

	NewInfrastructure := infrastructure.NewInfrastructure(db)
	eventUseCase := usecase.NewEventUsecase(NewInfrastructure)

	pg.RunMigrations()
	logger := slog.Default()

	Appconfig = &Config{
		Event:  eventUseCase,
		logger: logger,
	}
	
	return Appconfig
}