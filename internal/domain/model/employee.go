package model

import "time"

type EmployeeRole int

const (
	Specialist EmployeeRole = 1
	Leader     EmployeeRole = 2
	Manager    EmployeeRole = 3
)

type Employee struct {
	ID        int64
	Name      string
	Surname   string
	Phone     string
	Email     string
	Role      EmployeeRole
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewEmployee(name, surname, phone, email string, post EmployeeRole, now time.Time) *Employee {
	return &Employee{
		Name:      name,
		Surname:   surname,
		Phone:     phone,
		Email:     email,
		Role:      post,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (e *Employee) IsCanOrderCreate() bool {
	return e.Role == Manager
}

func (e *Employee) ChangeRole(newRole EmployeeRole) {
	e.Role = newRole
}
