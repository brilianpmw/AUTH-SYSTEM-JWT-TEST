package usecase

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

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
	c := gomock.NewController(t)
	mockTime := timeUtils.NewMockTimeUtility(c)
	type fields struct {
		repository *Repositories
		timeUtils  timeUtils.TimeUtility
	}
	type args struct {
		ctx            context.Context
		questStart     time.Time
		questDueDate   time.Time
		shopIDs        []int64
		missionGroupID int64
		totalMission   int
		timeNow        time.Time
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
					SubscriptionService: &MockSubscriptionService{
						dataGetPackageDetail: presentation.PackageDetail{
							PackageName:  "name",
							PackagePrice: 10000,
							PackageID:    100,
						},
					},
					MissionInfo: &MockMissionInfo{
						dataMissionGroupWithReward: presentation.MissionAndRewardData{
							QuestTitle: "test",
							RewardValue: sql.NullInt64{
								Valid: true,
								Int64: 100,
							},
						},
					},
					IrisReward: &MockIrisReward{},
				},
				timeUtils: mockTime,
			},
			args: args{
				ctx:            context.Background(),
				questStart:     time.Date(2023, 5, 8, 1, 1, 1, 1, time.Local),
				questDueDate:   time.Date(2024, 5, 8, 1, 1, 1, 1, time.Local),
				shopIDs:        []int64{1},
				missionGroupID: 1,
				totalMission:   2,
				timeNow:        time.Date(2023, 5, 8, 1, 1, 1, 1, time.Local),
			},
			wantErr: false,
		},
		{
			name:   "err SubscriptionService",
			active: true,
			fields: fields{
				repository: &Repositories{
					SubscriptionService: &MockSubscriptionService{
						errGetPackageDetail: errors.New("any"),
					},
					MissionInfo: &MockMissionInfo{
						dataMissionGroupWithReward: presentation.MissionAndRewardData{
							QuestTitle: "test",
							RewardValue: sql.NullInt64{
								Valid: true,
								Int64: 123,
							},
						},
					},
					IrisReward: &MockIrisReward{},
				},
				timeUtils: mockTime,
			},
			args: args{
				ctx:            context.Background(),
				questStart:     time.Date(2023, 5, 8, 1, 1, 1, 1, time.Local),
				questDueDate:   time.Date(2024, 5, 8, 1, 1, 1, 1, time.Local),
				shopIDs:        []int64{1},
				missionGroupID: 1,
				totalMission:   2,
				timeNow:        time.Date(2023, 5, 8, 1, 1, 1, 1, time.Local),
			},
			wantErr: true,
		},
		{
			name:   "error getMissionInfo",
			active: true,
			fields: fields{
				repository: &Repositories{
					SubscriptionService: &MockSubscriptionService{},
					MissionInfo: &MockMissionInfo{
						errMissionGroupWithReward: errors.New("any"),
					},
					IrisReward: &MockIrisReward{},
				},
				timeUtils: mockTime,
			},
			args: args{
				ctx:            context.Background(),
				questStart:     time.Date(2023, 5, 8, 1, 1, 1, 1, time.Local),
				questDueDate:   time.Date(2024, 5, 8, 1, 1, 1, 1, time.Local),
				shopIDs:        []int64{1},
				missionGroupID: 1,
				totalMission:   2,
				timeNow:        time.Date(2023, 5, 8, 1, 1, 1, 1, time.Local),
			},
			wantErr: true,
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
