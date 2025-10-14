package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/kennedymwaniki/events-api/internal/domain"
)


func (p *PGDBInstance) CreateEvent(ctx context.Context, e *domain.Event) error {
	qs := `INSERT INTO events (title, description, location, date, time, capacity) values(@title, @description, @location, @date, @time, @capacity)`
	_, err := p.Exec(ctx, qs, pgx.StrictNamedArgs{
		"title":        e.Title,
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