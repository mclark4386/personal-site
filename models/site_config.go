package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type SiteConfig struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Key       string    `json:"key" db:"key"`
	Value     string    `json:"value" db:"value"`
}

// String is not required by pop and may be deleted
func (s SiteConfig) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// SiteConfigs is not required by pop and may be deleted
type SiteConfigs []SiteConfig

// String is not required by pop and may be deleted
func (s SiteConfigs) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *SiteConfig) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: s.Key, Name: "Key"},
		&validators.StringIsPresent{Field: s.Value, Name: "Value"},
		// check to see if the key is already taken:
		&validators.FuncValidator{
			Field:   s.Key,
			Name:    "Key",
			Message: "%s is already taken",
			Fn: func() bool {
				var b bool
				q := tx.Where("key = ?", s.Key)
				if s.ID != uuid.Nil {
					q = q.Where("id != ?", s.ID)
				}
				b, err = q.Exists(s)
				if err != nil {
					return false
				}
				return !b
			},
		},
	), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *SiteConfig) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *SiteConfig) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
