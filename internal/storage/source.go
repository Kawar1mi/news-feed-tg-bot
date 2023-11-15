package storage

import (
	"context"
	"time"

	"github.com/Kawar1mi/news-feed-tg-bot/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type SourcePostgresStorage struct {
	db *sqlx.DB
}

func NewSourcePostgresStorage(db *sqlx.DB) *SourcePostgresStorage {
	return &SourcePostgresStorage{
		db: db,
	}
}

func (s *SourcePostgresStorage) Sources(ctx context.Context) ([]model.Source, error) {
	var sources []dbSource

	if err := s.db.GetContext(ctx, &sources, "SELECT * FROM sources"); err != nil {
		return nil, err
	}

	return lo.Map(sources, func(source dbSource, _ int) model.Source {
		return model.Source(source)
	}), nil
}

func (s *SourcePostgresStorage) SourceByID(ctx context.Context, id int64) (*model.Source, error) {
	var source dbSource

	if err := s.db.GetContext(ctx, &source, "SELECT * FROM sources WHERE id = $1", id); err != nil {
		return nil, err
	}

	return (*model.Source)(&source), nil
}

func (s *SourcePostgresStorage) Add(ctx context.Context, source model.Source) (int64, error) {
	s.db.QueryRowxContext()

}

func (s *SourcePostgresStorage) Delete(ctx context.Context, id int64) error {
	s.db.ExecContext()
}

type dbSource struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	FeedURL   string    `db:"feed_url"`
	Priority  int       `db:"priority"`
	CreatedAt time.Time `db:"created_at"`
}
