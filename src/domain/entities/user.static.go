package entities

import (
	"errors"
	"github.com/dgryski/trifles/uuid"
	"github.com/lvdigitalpro/back/src/domain/utils"
	"strings"
)

func (r *User) Validate() bool {

	if r.Role == RoleUser {

		if r.Ir == "" {
			return false
		}

		irValid := utils.IsCPF(r.Ir)

		if !irValid {
			return false
		}

	}

	if r.Role == RoleEnterprise {

		if r.Nrle == nil || r.Ir == "" {
			return false
		}

		irValid := utils.IsCPF(r.Ir)
		nrleValid := utils.IsCNPJ(*r.Nrle)

		if !irValid {
			return false
		}

		if !nrleValid {
			return false
		}

	}

	return true
}

func (r *User) AssignUser(user User) {
	r.ID = user.ID
	r.Role = user.Role
	r.Name = user.Name
	r.Lastname = user.Lastname
	r.Ir = user.Ir
	r.Email = user.Email
	r.Password = user.Password
	r.EnterpriseName = user.EnterpriseName
	r.Nrle = user.Nrle
	r.CreatedAt = user.CreatedAt
	r.UpdatedAt = user.UpdatedAt
	r.Projects = user.Projects
}

func (r *User) NewUser(role Role, name string, lastname string, ir string, email string, password string, passwordConfirmation string, enterpriseName *string, nrle *string) (*User, error) {

	if password != passwordConfirmation {
		return nil, errors.New("passwords do not match")
	}

	r.ID = uuid.UUIDv4()
	r.Role = role
	r.Name = strings.TrimSpace(name)
	r.Lastname = strings.TrimSpace(lastname)
	r.Ir = strings.TrimSpace(ir)
	r.Email = strings.TrimSpace(email)
	r.Password = strings.TrimSpace(password)

	if role == RoleEnterprise {
		r.EnterpriseName = enterpriseName
		r.Nrle = nrle
	}

	validate := r.Validate()
	if !validate {
		return nil, errors.New("invalid user")
	}

	return r, nil
}
