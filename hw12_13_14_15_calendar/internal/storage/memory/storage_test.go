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
			NotificationTime: time.Time{},
		}
		emptyList, err := st.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(emptyList))

		assert.NoError(t, st.Add(ctx, ev))

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
			NotificationTime: time.Time{},
		}
		assert.NoError(t, st.Add(ctx, ev))

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
			NotificationTime: time.Time{},
		}
		assert.NoError(t, st.Add(ctx, ev))

		f, err := st.Find(ctx, ev.ID)
		assert.NoError(t, err)
		assert.Equal(t, ev.Title, f.Title)

		updTitle := "upd title"
		ev.Title = updTitle
		assert.NoError(t, st.Update(ctx, ev))

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
			NotificationTime: time.Time{},
		}
		ev1 := storage.Event{
			ID:               uuid.New(),
			Title:            "some title",
			Start:            time.Time{},
			End:              time.Time{},
			Description:      "some description",
			UserID:           uuid.New(),
			NotificationTime: time.Time{},
		}
		ev2 := storage.Event{
			ID:               uuid.New(),
			Title:            "some title",
			Start:            time.Time{},
			End:              time.Time{},
			Description:      "some description",
			UserID:           uuid.New(),
			NotificationTime: time.Time{},
		}
		emptyList, err := st.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(emptyList))

		assert.NoError(t, st.Add(ctx, ev))
		assert.NoError(t, st.Add(ctx, ev1))
		assert.NoError(t, st.Add(ctx, ev2))

		resList, err := st.List(ctx)
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
			NotificationTime: time.Time{},
		}

		assert.NoError(t, st.Add(ctx, ev))
		l, err := st.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(l))

		assert.NoError(t, st.Delete(ctx, ev.ID))
		l, err = st.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(l))
	})
}
