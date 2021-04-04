package domain

import (
	"time"
	"database/sql/driver"
	"github.com/google/uuid"

	"github.com/mysa-fika/go/infrastructure/storage"
	"github.com/mysa-fika/go/domain/vocabulary"
)

type Repository interface {
	Add(User) error
}

type Lock struct {
	Code   uuid.UUID             `json:"code"`
	Reason vocabulary.LockReason `json:"reason"`
}

func (l *Lock) Scan(src interface{}) error {
	return storage.ScanJSONB(l, src)
}

func (l Lock) Value() (driver.Value, error) {
	return storage.JSONBValue(l)
}

type Metadata struct{}

func (m *Metadata) Scan(src interface{}) error {
	return storage.ScanJSONB(m, src)
}

func (m Metadata) Value() (driver.Value, error) {
	return storage.JSONBValue(m)
}

type User struct {
	ID           uuid.UUID        `db:"id"`
	CreatedTime  time.Time        `db:"created_time"`
	ModifiedTime time.Time        `db:"modified_time"`
	Realm        vocabulary.Realm `db:"realm"`
	Email        string           `db:"email"`
	Password     string           `db:"password"`
	Lock         *Lock            `db:"lock,omitempty"`
	Metadata     Metadata         `db:"metadata"`
}
