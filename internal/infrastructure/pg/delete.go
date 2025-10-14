package pg

import (
	"context"
)

func (p *PGDBInstance) DeleteEvent(ctx context.Context, id int) error {
	qs := `DELETE FROM events WHERE id = $1`
	
	_, err := p.Exec(ctx, qs, id)
	if err != nil {
		return err
	}
	return nil
}