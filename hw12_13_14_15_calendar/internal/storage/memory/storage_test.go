package memorystorage

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	ctx := context.Background()
	t.Run("add", func(t *testing.T) {
		st := New()
		ev := storage.Event{
			ID:               uuid.New(),
			Title:            "some title",
			Start:            time.Time{},
			End:              time.Time{},
			Description:      "some description",
			UserID:           uuid.New(),
			NotificationTime: time.Duration(0),
		}
		emptyList, err := st.ListByInterval(ctx, time.Now(), time.Now())
		assert.NoError(t, err)
		assert.Equal(t, 0, len(emptyList))

		_, e := st.Add(ctx, ev)
		assert.NoError(t, e)

		evFound, err := st.Find(ctx, ev.ID)
		assert.NoError(t, err)
		assert.Equal(t, ev, *evFound)
	})

	t.Run("find", func(t *testing.T) {
		st := New()
		ev := storage.Event{
			ID:               uuid.New(),
			Title:            "some title",
			Start:            time.Time{},
			End:              time.Time{},
			Description:      "some description",
			UserID:           uuid.New(),
			NotificationTime: time.Duration(0),
		}
		_, e := st.Add(ctx, ev)
		assert.NoError(t, e)

		evFound, err := st.Find(ctx, ev.ID)
		assert.NoError(t, err)
		assert.Equal(t, ev, *evFound)
	})

	t.Run("update", func(t *testing.T) {
		st := New()
		ev := storage.Event{
			ID:               uuid.New(),
			Title:            "some title",
			Start:            time.Time{},
			End:              time.Time{},
			Description:      "some description",
			UserID:           uuid.New(),
			NotificationTime: time.Duration(0),
		}
		_, e := st.Add(ctx, ev)
		assert.NoError(t, e)

		f, err := st.Find(ctx, ev.ID)
		assert.NoError(t, err)
		assert.Equal(t, ev.Title, f.Title)

		updTitle := "upd title"
		ev.Title = updTitle

		_, e = st.Update(ctx, ev)
		assert.NoError(t, e)

		evFound, err := st.Find(ctx, ev.ID)
		assert.NoError(t, err)
		assert.Equal(t, ev, *evFound)
	})

	t.Run("list", func(t *testing.T) {
		st := New()
		ev := storage.Event{
			ID:               uuid.New(),
			Title:            "some title",
			Start:            time.Time{},
			End:              time.Time{},
			Description:      "some description",
			UserID:           uuid.New(),
			NotificationTime: time.Duration(0),
		}
		ev1 := storage.Event{
			ID:               uuid.New(),
			Title:            "some title",
			Start:            time.Time{},
			End:              time.Time{},
			Description:      "some description",
			UserID:           uuid.New(),
			NotificationTime: time.Duration(0),
		}
		ev2 := storage.Event{
			ID:               uuid.New(),
			Title:            "some title",
			Start:            time.Time{},
			End:              time.Time{},
			Description:      "some description",
			UserID:           uuid.New(),
			NotificationTime: time.Duration(0),
		}
		emptyList, err := st.ListByInterval(ctx, time.Now(), time.Now())
		assert.NoError(t, err)
		assert.Equal(t, 0, len(emptyList))

		_, e := st.Add(ctx, ev)
		assert.NoError(t, e)
		_, e = st.Add(ctx, ev1)
		assert.NoError(t, e)
		_, e = st.Add(ctx, ev2)
		assert.NoError(t, e)

		resList, err := st.ListByInterval(ctx, time.Now(), time.Now())
		assert.NoError(t, err)
		assert.Equal(t, 3, len(resList))
	})

	t.Run("delete", func(t *testing.T) {
		st := New()
		ev := storage.Event{
			ID:               uuid.New(),
			Title:            "some title",
			Start:            time.Time{},
			End:              time.Time{},
			Description:      "some description",
			UserID:           uuid.New(),
			NotificationTime: time.Duration(0),
		}

		_, e := st.Add(ctx, ev)
		assert.NoError(t, e)
		l, err := st.ListByInterval(ctx, time.Now(), time.Now())
		assert.NoError(t, err)
		assert.Equal(t, 1, len(l))

		assert.NoError(t, st.Delete(ctx, ev.ID))
		l, err = st.ListByInterval(ctx, time.Now(), time.Now())
		assert.NoError(t, err)
		assert.Equal(t, 0, len(l))
	})
}
