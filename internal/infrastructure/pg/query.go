package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/kennedymwaniki/events-api/internal/domain"
)

func (p *PGDBInstance) GetEvent(ctx context.Context, id string) (*domain.Event, error) {
	qs := `SELECT id, title, description, location, date, time, capacity, created_at, updated_at FROM events WHERE id = $1`
	
	row := p.QueryRow(ctx, qs, id)
	
	result := &domain.Event{}
	err := row.Scan(&result.Id, &result.Title, &result.Description, &result.Location, &result.Date, &result.Time, &result.Capacity, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	
	return result, nil
}

func (p *PGDBInstance) GetAllEvents(ctx context.Context) ([]*domain.Event, error) {
	qs := `SELECT id, title, description, location, date, time, capacity, created_at, updated_at FROM events`

	rows, err := p.Query(ctx, qs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*domain.Event
	for rows.Next() {
		result := &domain.Event{}

		err := rows.Scan(&result.Id, &result.Title, &result.Description, &result.Location, &result.Date, &result.Time, &result.Capacity, &result.CreatedAt, &result.UpdatedAt)
		if err != nil {
			return nil, err
		}

		events = append(events, result)
	}

	return events, nil
}

func (p *PGDBInstance) UpdateEvent(ctx context.Context, e *domain.Event, id string) error {
	qs := `UPDATE events SET title = @title, description = @description, location = @location, date = @date, time = @time, capacity = @capacity, updated_at = NOW() WHERE id = @id`
	
	_, err := p.Exec(ctx, qs, pgx.StrictNamedArgs{
		"id":          e.Id,
		"title":       e.Title,
		"description": e.Description,
		"location":    e.Location,
		"date":        e.Date,
		"time":        e.Time,
		"capacity":    e.Capacity,
	})
	if err != nil {
		return err
	}
	return nil
}

