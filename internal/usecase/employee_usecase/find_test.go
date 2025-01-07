package employee_usecase

import (
	"errors"
	"testing"

	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/employee_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestFindUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	employee := &model.Employee{
		ID:   1,
		Name: "Name",
	}

	type fields struct {
		employeeRepo *mocks.EmployeeRepo
		timer        *mocks.Timer
	}

	type args struct {
		req FindEmployeeReq
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
				req: FindEmployeeReq{
					ID: 1,
				},
			},
			want: &model.Employee{
				ID:   1,
				Name: "Name",
			},
			before: func(f fields, args args) {
				f.employeeRepo.EXPECT().FindEmployee(args.req.ID).Return(employee, nil)
			},
		},
		{
			name: "error on find",
			args: args{
				req: FindEmployeeReq{
					ID: 2,
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

			c, err := uc.FindEmployee(tt.args.req)

			if tt.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
			a.Equal(tt.want, c)
		})
	}
}
