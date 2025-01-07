package employee_usecase

import (
	"errors"
	"testing"

	"warehouse_project/internal/usecase/employee_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestDeleteUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	type fields struct {
		employeeRepo *mocks.EmployeeRepo
		timer        *mocks.Timer
	}

	type args struct {
		req DeleteEmployeeReq
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				req: DeleteEmployeeReq{
					ID: 1,
				},
			},
			before: func(f fields, args args) {
				f.employeeRepo.EXPECT().DeleteEmployee(args.req.ID).Return(nil)
			},
		},
		{
			name: "error on find",
			args: args{
				req: DeleteEmployeeReq{
					ID: 50,
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.employeeRepo.EXPECT().DeleteEmployee(args.req.ID).Return(errTest)
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

			err := uc.DeleteEmployee(tt.args.req)

			if tt.wantErr {
				a.Error(err)
				return
			}

			a.NoError(err)
		})
	}
}
