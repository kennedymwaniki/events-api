package pg

import (
	"context"

	"github.com/kennedymwaniki/events-api/internal/domain"
)

type EventsRepository interface {
	CreateEvent(ctx context.Context, e *domain.Event) error
	GetEvent(ctx context.Context, id string) (*domain.Event, error)
	GetAllEvents(ctx context.Context) ([]*domain.Event, error)
	UpdateEvent(ctx context.Context, e *domain.Event, id string) error
	DeleteEvent(ctx context.Context, id int) error
}