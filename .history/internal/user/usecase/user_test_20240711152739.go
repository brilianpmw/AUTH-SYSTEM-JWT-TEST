package usecase

import (
	"context"
	"testing"

	"github.com/brilianpmw/synapsis/internal/pkg/config"
	"github.com/brilianpmw/synapsis/presentation"
)

type MockUserUsecase struct {
	dataGetUserDataByUserName  presentation.User
	errorGetUserDataByUserName error
}

func (m *MockUserUsecase) GetUserDataByUserName(ctx context.Context, username string) (presentation.User, error) {
	return m.dataGetUserDataByUserName, m.errorGetUserDataByUserName
}

func TestUsecase_DoLogin(t *testing.T) {
	type fields struct {
		repository *Repositories
	}
	type args struct {
		request presentation.LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		active  bool
	}{

		{
			name:   "success",
			active: true,
			fields: fields{
				repository: &Repositories{
					User: &MockUserUsecase{
						dataGetUserDataByUserName: presentation.User{
							Username: "Brilian",
							Gender:   "Man",
							Password: "hashedpw",
						},
					},
				},
			},
			args: args{
				request: presentation.LoginRequest{
					Username: "Brilian",
					Password: "hashedpw",
				},
			},
			wantErr: false,
		},
		{
			name:   "empty data",
			active: true,
			fields: fields{
				repository: &Repositories{
					User: &MockUserUsecase{
						dataGetUserDataByUserName: presentation.User{
							Username: "Brilian",
							Gender:   "Man",
							Password: "hashedpw",
						},
					},
				},
			},
			args: args{
				request: presentation.LoginRequest{
					Username: "Brilian",
					Password: "hashedpw",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &Usecase{
				repository: tt.fields.repository,
			}
			config.InitTestingConfig()
			if _, err := uc.DoLogin(context.Background(), tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Usecase.DoLogin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
