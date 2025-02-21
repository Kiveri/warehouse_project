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
	Phone     string
	Email     string
	Role      EmployeeRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewEmployee(name, phone, email string, role EmployeeRole, now time.Time) *Employee {
	return &Employee{
		Name:      name,
		Phone:     phone,
		Email:     email,
		Role:      role,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (e *Employee) IsCanOrderCreate() bool {
	return e.Role == Manager
}

func (e *Employee) ChangeRole(newRole EmployeeRole, now time.Time) {
	e.Role = newRole
	e.UpdatedAt = now
}
