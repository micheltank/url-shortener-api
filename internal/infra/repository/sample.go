package repository

import (
	"github.com/micheltank/url-shortener-api/internal/domain"
	"github.com/micheltank/url-shortener-api/internal/infra/cassandra"
)

// Repository is the feed service's wrapper around database access
type Repository struct {
	Wrapper
}

// NewRepository constructs a new repository
func NewRepository(cassandra *cassandra.Client) *Repository {
	return &Repository{
		Wrapper{
			Cassandra: cassandra,
		},
	}
}

// Get retrieves the feed of tweets for a particular user
func (r *Repository) Get(hash string) (url domain.Url, err error) {
	var longUrl string

	query := r.Query("SELECT hash, email, refresh_token FROM users WHERE username = ?", hash)

	err := query.Scan(&hash, &longUrl)

	if err != nil {
		return url, err
	}

	return domain.NewShortenedUrl(hash, longUrl), nil
}

func (r *Repository) Create(hash, longUrl string) error {
	query := r.Query(`
		INSERT INTO feed_items
			(username, tweet_created_at, tweet_content, tweet_id, tweet_username)
		VALUES
			(?, ?, ?, ?, ?)
	`,
		followerUsername, item.CreatedAt, item.Content, item.ID.String(), item.Username,
	)

	err := query.Exec()

	if err != nil {
		return err
	}

	return nil
}
