package event

import (
	"context"

	"github.com/kennedymwaniki/events-api/internal/domain"
	"github.com/kennedymwaniki/events-api/internal/infrastructure"
)

type EventUseCase struct {
	r infrastructure.Infrastructure
}

func NewEventUsecase(r infrastructure.Infrastructure) *EventUseCase {
	return &EventUseCase{
		r: r,
	}
}

func (q *EventUseCase) CreateEvent(ctx context.Context, e *domain.Event) error {
	return q.r.EventsRepository.CreateEvent(ctx, e)
}
	
func (q *EventUseCase) GetEvent(ctx context.Context, id string) (*domain.Event, error) {
	return q.r.EventsRepository.GetEvent(ctx, id)
}

func (q *EventUseCase) GetAllEvents(ctx context.Context) ([]*domain.Event, error) {
	return q.r.EventsRepository.GetAllEvents(ctx)
}


func (q *EventUseCase) UpdateEvent(ctx context.Context, e *domain.Event, id string) error {
	return q.r.EventsRepository.UpdateEvent(ctx, e, id)
}

func (q *EventUseCase) DeleteEvent(ctx context.Context, id int) error {
	return q.r.EventsRepository.DeleteEvent(ctx, id)
}