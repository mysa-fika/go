package memory

import (
	"context"
	"github.com/mysa-fika/go/authentication-service/pkg/domain"
	"github.com/mysa-fika/go/infrastructure/log"
	"sync"
)

type Repository struct {
	DB     sync.Map
	Logger *log.Logger
}

func (r *Repository) Add(u domain.User) error {
	r.DB.Store(u.ID, u)
	return nil
}

func NewRepository(ctx context.Context, _ map[string]interface{}, logger *log.Logger) (*Repository, error) {
	return &Repository{
		DB:     sync.Map{},
		Logger: logger,
	}, nil
}
