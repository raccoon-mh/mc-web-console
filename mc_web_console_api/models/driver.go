package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Driver is used by pop to map your drivers database table to your go code.
type Driver struct {
	ID          uuid.UUID `json:"id" db:"id"`
	DriverName  string    `json:"driver_name" db:"driver_name"`
	LibFileName string    `json:"lib_file_name" db:"lib_file_name"`

	ProviderID string         `json:"provider_id" db:"provider_id"`
	Provider   *CloudProvider `belongs_to:"cloud_provider"`

	//CloudConnection *CloudConnection `belongs_to:"cloud_connection"`
	CloudConnections []CloudConnection `has_many:"cloud_connection"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Drivers is not required by pop and may be deleted
type Drivers []Driver

func (d *Driver) Create(tx *pop.Connection) (*validate.Errors, error) {
	log.Println("Create  ", d)
	return tx.ValidateAndCreate(d)
}

func (d *Driver) Update(tx *pop.Connection) (*validate.Errors, error) {
	log.Println("Update  ", d)
	return tx.ValidateAndUpdate(d)
}

// String is not required by pop and may be deleted
func (d Driver) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// String is not required by pop and may be deleted
func (d Drivers) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Driver) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	//return validate.NewErrors(), nil
	return nil, err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Driver) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	// return validate.Validate(
	// 	&validators.StringIsPresent{Field: d.DriverName, Name: "DriverName"},
	// ), err
	//return tx.ValidateAndCreate(d)
	return nil, err
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Driver) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	// return validate.NewErrors(), nil
	// return tx.ValidateAndUpdate(d)
	return nil, err
}
