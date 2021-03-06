package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Issue struct {
	ID               uuid.UUID     `json:"id" db:"id"`
	CreatedAt        time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at" db:"updated_at"`
        GithubUpdatedAt  time.Time     `json:"github_updated_at" db:"github_updated_at"`
	Title            nulls.String  `json:"title" db:"title"`
	ExperienceNeeded nulls.String  `json:"experience_needed" db:"experience_needed"`
	ExpectedTime     nulls.String  `json:"expected_time" db:"expected_time"`
	Language         nulls.String  `json:"language" db:"language"`
	TechStack        nulls.String  `json:"tech_stack" db:"tech_stack"`
	GithubID         int           `json:"github_id" db:"github_id"`
	URL              string        `json:"url" db:"url"`
	Body             nulls.String  `json:"body" db:"body"`
	Type             nulls.String  `json:"type" db:"type"`
	Repository       Repository    `json:"repository" db:"-" belongs_to:"repository"`
	RepositoryID     uuid.UUID     `json:"repository_id" db:"repository_id" `
	Project          Project       `json:"project" db:"-" belongs_to:"project"`
	ProjectID        uuid.UUID     `json:"project_id" db:"project_id" `
	Number           int           `json:"number" db:"number"`
	Closed           bool          `json:"-" db:"closed"`
	Labels           slices.String `json:"labels" db:"labels"`
}

// String is not required by pop and may be deleted
func (i Issue) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Issues is not required by pop and may be deleted
type Issues []Issue

// String is not required by pop and may be deleted
func (i Issues) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (i *Issue) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: i.GithubID, Name: "GithubID"},
		&validators.IntIsPresent{Field: i.Number, Name: "Number"},
		&validators.StringIsPresent{Field: i.URL, Name: "URL"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (i *Issue) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (i *Issue) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
