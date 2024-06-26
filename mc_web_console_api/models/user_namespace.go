package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// UserNamespace is used by pop to map your user_namespaces database table to your go code.
type UserNamespace struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	NamespaceID string     `db:"ns_id"`
	Namespace   *Namespace `belongs_to:"namespaces"`
	UserID      uuid.UUID  `json:"user_id" db:"user_id"`
	User        *User      `belongs_to:"users"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (u UserNamespace) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// UserNamespaces is not required by pop and may be deleted
type UserNamespaces []UserNamespace

// String is not required by pop and may be deleted
func (u UserNamespaces) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *UserNamespace) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *UserNamespace) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *UserNamespace) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (u *UserNamespace) Create(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(u)
}
