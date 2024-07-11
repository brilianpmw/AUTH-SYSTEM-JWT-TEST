package usecase

import (
	"context"
	"testing"

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
				repository: &Repositories{},
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &Usecase{
				repository: tt.fields.repository,
				timeUtils:  tt.fields.timeUtils,
			}
			config.InitTestingConfig(&config.Configuration{
				FeatureFlags: config.FFConfig{
					ActivateSendQuestStartDataToIris: tt.active,
				},
			})
			if err := uc.publishIrisEventStart(tt.args.ctx, tt.args.questStart, tt.args.questDueDate, tt.args.shopIDs, tt.args.missionGroupID, tt.args.totalMission, tt.args.timeNow); (err != nil) != tt.wantErr {
				t.Errorf("Usecase.publishIrisEventStart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
