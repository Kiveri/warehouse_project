package employee_usecase

import (
	"errors"
	"testing"
	"time"

	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/employee_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestChangeRoleUseCase(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("test error")

	type fields struct {
		employeeRepo *mocks.EmployeeRepo
		timer        *mocks.Timer
	}

	type args struct {
		req UpdateEmployeeReq
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Employee
		wantErr bool
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				req: UpdateEmployeeReq{
					ID:   1,
					Role: model.Specialist,
				},
			},
			want: &model.Employee{
				ID:        1,
				Role:      model.Specialist,
				UpdatedAt: now,
			},
			before: func(f fields, args args) {
				employee := &model.Employee{
					ID:   1,
					Role: model.Leader,
				}

				f.employeeRepo.EXPECT().FindEmployee(args.req.ID).Return(employee, nil)

				f.timer.EXPECT().Now().Return(now)
				employee.Role = args.req.Role
				employee.UpdatedAt = now

				f.employeeRepo.EXPECT().UpdateEmployee(employee).Return(employee, nil)
			},
		},
		{
			name: "error on update",
			args: args{
				req: UpdateEmployeeReq{
					ID:   1,
					Role: model.Specialist,
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				employee := &model.Employee{
					ID:   1,
					Role: model.Leader,
				}

				f.employeeRepo.EXPECT().FindEmployee(args.req.ID).Return(employee, nil)

				f.timer.EXPECT().Now().Return(now)
				employee.Role = args.req.Role
				employee.UpdatedAt = now

				f.employeeRepo.EXPECT().UpdateEmployee(employee).Return(nil, errTest)
			},
		},
		{
			name: "error on find",
			args: args{
				req: UpdateEmployeeReq{
					ID:   1,
					Role: model.Specialist,
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.employeeRepo.EXPECT().FindEmployee(args.req.ID).Return(nil, errTest)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			f := fields{
				employeeRepo: mocks.NewEmployeeRepo(t),
				timer:        mocks.NewTimer(t),
			}
			tt.before(f, tt.args)

			uc := NewEmployeeUseCase(f.employeeRepo, f.timer)

			e, err := uc.UpdateEmployee(tt.args.req)

			if tt.wantErr {
				a.Error(err)

				return
			}

			a.NoError(err)
			a.Equal(tt.want, e)
		})
	}
}
