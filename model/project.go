package model

import (
	"net/http"

	"github.com/anthonynsimon/parrot/errors"
)

var (
	ErrInvalidProjectName = errors.New(
		http.StatusBadRequest,
		"InvalidProjectName",
		"invalid field project name")
)

type ProjectStorer interface {
	GetProjects() ([]Project, error)
	GetProject(int) (*Project, error)
	CreateProject(*Project) (Project, error)
	UpdateProject(*Project) error
	DeleteProject(int) (int, error)
}

type ProjectLocaleStorer interface {
	GetProjectLocale(projID, localeID int) (*Locale, error)
	FindProjectLocales(projID int, localeIdents ...string) ([]Locale, error)
}

type ProjectUserStorer interface {
	GetProjectUsers(projID int) ([]User, error)
	GetUserProjects(userID int) ([]Project, error)
	GetProjectUserRoles(projID int) ([]ProjectUser, error)
	AssignProjectUser(ProjectUser) error
	RevokeProjectUser(ProjectUser) error
	UpdateProjectUser(ProjectUser) (*ProjectUser, error)
}

type Project struct {
	ID   int      `db:"id" json:"id"`
	Name string   `db:"name" json:"name"`
	Keys []string `db:"keys" json:"keys"`
}

type ProjectUser struct {
	ProjectID int    `json:"project_id"`
	UserID    int    `json:"user_id"`
	Role      string `json:"role"`
}

func (p *Project) SanitizeKeys() {
	var sk []string
	for _, key := range p.Keys {
		if key == "" {
			continue
		}
		sk = append(sk, key)
	}

	p.Keys = sk
}

func (p *Project) Validate() error {
	var errs []errors.Error
	if !HasMinLength(p.Name, 1) {
		errs = append(errs, *ErrInvalidProjectName)
	}
	if errs != nil {
		return &errors.MultiError{Errors: errs}
	}
	return nil
}
