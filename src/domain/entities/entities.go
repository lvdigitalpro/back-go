// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package entities

import (
	"fmt"
	"io"
	"strconv"
)

type Project struct {
	ID          string   `json:"id"`
	Type        Type     `json:"type"`
	Owner       *User    `json:"owner"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tasks       []*Tasks `json:"tasks,omitempty"`
	Status      Status   `json:"status"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
}

type Tasks struct {
	ID          string   `json:"id"`
	Type        TypeTask `json:"type"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	Owner       *User    `json:"owner"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}

type User struct {
	ID             string   `json:"id"`
	Role           Role     `json:"role"`
	Name           string   `json:"name"`
	Lastname       string   `json:"lastname"`
	Ir             string   `json:"ir"`
	Email          string   `json:"email"`
	Password       string   `json:"password"`
	EnterpriseName *string  `json:"enterpriseName,omitempty"`
	Nrle           *string  `json:"nrle,omitempty"`
	Projects       []string `json:"projects,omitempty"`
	CreatedAt      string   `json:"createdAt"`
	UpdatedAt      *string  `json:"updatedAt,omitempty"`
}

type Role string

const (
	RoleAdmin      Role = "ADMIN"
	RoleUser       Role = "USER"
	RoleEnterprise Role = "ENTERPRISE"
)

var AllRole = []Role{
	RoleAdmin,
	RoleUser,
	RoleEnterprise,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser, RoleEnterprise:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Status string

const (
	StatusQueue      Status = "QUEUE"
	StatusInReview   Status = "IN_REVIEW"
	StatusInProgress Status = "IN_PROGRESS"
	StatusFinished   Status = "FINISHED"
)

var AllStatus = []Status{
	StatusQueue,
	StatusInReview,
	StatusInProgress,
	StatusFinished,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusQueue, StatusInReview, StatusInProgress, StatusFinished:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Type string

const (
	TypeWeb     Type = "WEB"
	TypeMobile  Type = "MOBILE"
	TypeDesktop Type = "DESKTOP"
	TypeSystem  Type = "SYSTEM"
	TypeUIUx    Type = "UI_UX"
	TypeOther   Type = "OTHER"
)

var AllType = []Type{
	TypeWeb,
	TypeMobile,
	TypeDesktop,
	TypeSystem,
	TypeUIUx,
	TypeOther,
}

func (e Type) IsValid() bool {
	switch e {
	case TypeWeb, TypeMobile, TypeDesktop, TypeSystem, TypeUIUx, TypeOther:
		return true
	}
	return false
}

func (e Type) String() string {
	return string(e)
}

func (e *Type) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Type(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}

func (e Type) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TypeTask string

const (
	TypeTaskSearch      TypeTask = "SEARCH"
	TypeTaskDesign      TypeTask = "DESIGN"
	TypeTaskDevelopment TypeTask = "DEVELOPMENT"
	TypeTaskTest        TypeTask = "TEST"
	TypeTaskOther       TypeTask = "OTHER"
)

var AllTypeTask = []TypeTask{
	TypeTaskSearch,
	TypeTaskDesign,
	TypeTaskDevelopment,
	TypeTaskTest,
	TypeTaskOther,
}

func (e TypeTask) IsValid() bool {
	switch e {
	case TypeTaskSearch, TypeTaskDesign, TypeTaskDevelopment, TypeTaskTest, TypeTaskOther:
		return true
	}
	return false
}

func (e TypeTask) String() string {
	return string(e)
}

func (e *TypeTask) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TypeTask(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TypeTask", str)
	}
	return nil
}

func (e TypeTask) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}