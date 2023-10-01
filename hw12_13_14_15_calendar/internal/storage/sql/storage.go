package sqlstorage

import (
	"context"

	"github.com/google/uuid"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	client ClientInterface
}

func New(client ClientInterface) *Storage {
	return &Storage{client: client}
}

func (s *Storage) Add(ctx context.Context, e storage.Event) error {
	q := `INSERT INTO events (id, title, start, "end", description, user_id, notification_time)
			VALUES (:id, :title, :start, :end, :description, :user_id, :notification_time);`
	_, err := s.client.Connection().NamedExecContext(ctx, q, e)
	return err
}

func (s *Storage) Find(ctx context.Context, uuid uuid.UUID) (*storage.Event, error) {
	q := `SELECT id, title, start, "end", description, user_id, notification_time FROM events WHERE id=$1`
	var ev storage.Event
	if err := s.client.Connection().GetContext(ctx, &ev, q, uuid); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (s *Storage) Update(ctx context.Context, e storage.Event) error {
	q := `UPDATE events 
		SET title=:title,
		    start=:start,
		    "end"=:end,
		    description=:description,
		    user_id=:user_id,
		    notification_time=:notification_time
		    WHERE id=:id`
	res, err := s.client.Connection().NamedExecContext(ctx, q, e)
	if err != nil {
		return err
	}

	c, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if c == 0 {
		return storage.ErrUpdateNoAffectedRows
	}
	return nil
}

func (s *Storage) Delete(ctx context.Context, uuid uuid.UUID) error {
	q := "DELETE FROM events WHERE id=$1"
	if _, err := s.client.Connection().ExecContext(ctx, q, uuid); err != nil {
		return err
	}
	return nil
}

func (s *Storage) List(ctx context.Context) ([]storage.Event, error) {
	q := "SELECT * FROM events"
	var events []storage.Event
	if err := s.client.Connection().SelectContext(ctx, &events, q); err != nil {
		return nil, err
	}
	return events, nil
}
