package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/Kawar1mi/news-feed-tg-bot/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type ArticlePostgresStorage struct {
	db *sqlx.DB
}

func NewArticlePostgresStorage(db *sqlx.DB) *ArticlePostgresStorage {
	return &ArticlePostgresStorage{
		db: db,
	}
}

func (a *ArticlePostgresStorage) Store(ctx context.Context, article model.Article) error {
	_, err := a.db.ExecContext(
		ctx,
		`INSERT INTO articles (source_id, title, link, summary, published_at)
						VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`,
		article.SourceID,
		article.Title,
		article.Link,
		article.Summary,
		article.PublishedAt)

	if err != nil {
		return err
	}

	return nil
}

func (a *ArticlePostgresStorage) AllNotPosted(ctx context.Context, since time.Time, limit uint64) ([]model.Article, error) {
	var articles []dbArticle

	err := a.db.SelectContext(
		ctx,
		&articles,
		`SELECT 
			a.id,
			s.priority,
			s.id as source_id,
			a.title,
			a.link,
			a.summary,
			a.published_at,
			a.created_at,
			a.posted_at
		FROM articles a JOIN sources s ON a.source_id = s.id
		WHERE a.posted_at IS NULL AND a.published_at >= $1::timestamp
		ORDER BY a.created_at DESC, s.priority DESC LIMIT $2`,
		since.UTC().Format(time.RFC3339),
		limit)

	if err != nil {
		return nil, err
	}

	return lo.Map(articles, func(article dbArticle, _ int) model.Article {
		return model.Article{
			ID:          article.ID,
			SourceID:    article.SourceID,
			Title:       article.Title,
			Link:        article.Link,
			Summary:     article.Summary.String,
			PublishedAt: article.Published_at,
			CreatedAt:   article.CreatedAt,
		}
	}), nil
}

func (a *ArticlePostgresStorage) MarkAsPosted(ctx context.Context, article model.Article) error {
	_, err := a.db.ExecContext(
		ctx,
		"UPDATE articles SET posted_at = $1::timestamp WHERE id = $2",
		time.Now().UTC().Format(time.RFC3339),
		article.ID)

	if err != nil {
		return err
	}

	return nil
}

type dbArticle struct {
	ID             int64          `db:"id"`
	SourcePriority int64          `db:"priority"`
	SourceID       int64          `db:"source_id"`
	Title          string         `db:"title"`
	Link           string         `db:"link"`
	Summary        sql.NullString `db:"summary"`
	Published_at   time.Time      `db:"published_at"`
	CreatedAt      time.Time      `db:"created_at"`
	Posted_at      sql.NullTime   `db:"posted_at"`
}
