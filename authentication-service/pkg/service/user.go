package service

import (
	"time"

	"github.com/google/uuid"

	"github.com/mysa-fika/go/authentication-service/pkg/domain"
	"github.com/mysa-fika/go/infrastructure/log"
	"github.com/mysa-fika/go/domain/vocabulary"
)

type UserService struct {
	Repository domain.Repository
	Logger     *log.Logger
}

func (s UserService) CreateUser(email, password, realm string) error {
	// hashing password
	now := time.Now().UTC()
	userID := uuid.New()
	// create lock
	newLock := &domain.Lock{}

	return s.Repository.Add(domain.User{
		ID:           userID,
		CreatedTime:  now,
		ModifiedTime: now,
		Realm:        vocabulary.Realm(realm),
		Email:        email,
		Lock:         newLock,
		Metadata:     domain.Metadata{},
	})
}
