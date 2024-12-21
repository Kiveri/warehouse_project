package employee_usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/employee_usecase/mocks"
)

func TestCreateUseCase(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("test error")

	// зависимости которые нужны для теста
	type fields struct {
		employeeRepo *mocks.EmployeeRepo
		timer        *mocks.Timer
	}
	// данные для теста
	type args struct {
		req CreateEmployeeReq
	}
	// тесты
	tests := []struct {
		name    string
		args    args
		want    *model.Employee
		wantErr bool                      // если хотим ошибку
		before  func(f fields, args args) // замокать наши вызовы перед тестом
	}{
		{
			name: "success",
			args: args{
				req: CreateEmployeeReq{
					Name:  "Denis Popov",
					Phone: "89995398037",
					Email: "denis@gmail.com",
					Role:  model.Manager,
				},
			},
			want: &model.Employee{
				Name:      "Denis Popov",
				Phone:     "89995398037",
				Email:     "denis@gmail.com",
				Role:      model.Manager,
				CreatedAt: now,
				UpdatedAt: now,
			},
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)
				employee := model.NewEmployee(args.req.Name, args.req.Phone, args.req.Email, args.req.Role, now)
				f.employeeRepo.EXPECT().CreateEmployee(employee).Return(employee, nil)
			},
		},
		{
			name: "error on create",
			args: args{
				req: CreateEmployeeReq{
					Name:  "Denis Popov",
					Phone: "89995398037",
					Email: "denis@gmail.com",
					Role:  model.Manager,
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)
				employee := model.NewEmployee(args.req.Name, args.req.Phone, args.req.Email, args.req.Role, now)
				f.employeeRepo.EXPECT().CreateEmployee(employee).Return(nil, errTest)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			// создали зависимости
			f := fields{
				employeeRepo: mocks.NewEmployeeRepo(t),
				timer:        mocks.NewTimer(t),
			}
			tt.before(f, tt.args)

			uc := NewEmployeeUseCase(f.employeeRepo, f.timer)

			// выполнили
			employee, err := uc.CreateEmployee(tt.args.req)

			// проверяем результат
			if tt.wantErr {
				a.Error(err)

				return
			}
			a.NoError(err)
			a.Equal(tt.want, employee)
		})
	}
}
