package infrastructure

import (
	"github.com/kennedymwaniki/events-api/internal/infrastructure/pg"
)


type Infrastructure struct {
	EventsRepository pg.EventsRepository
}

func NewInfrastructure(repo pg.EventsRepository) Infrastructure {
	return Infrastructure{
		EventsRepository: repo,
	}
}