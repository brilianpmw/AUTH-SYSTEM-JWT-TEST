package http

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/brilianpmw/synapsis/internal/config"
	"github.com/brilianpmw/synapsis/presentation"
	"github.com/go-chi/chi"
)

type MockUsecases2 struct {
	dataDoLogin string
	errDoLogin  error
}

func (m *MockUsecases2) DoLogin(ctx context.Context, req presentation.LoginRequest) (string, error) {
	return m.dataDoLogin, m.errDoLogin
}

func TestHttpHandler_HandleCreateMission(t *testing.T) {

	type fields struct {
		usecase presentation.IUserUC
	}
	tests := []struct {
		name        string
		fields      fields
		requestbody []byte
		want        []byte
		wantStatus  int
		timeUtils   timeUtils.TimeUtility
		expect      []*gomock.Call
	}{
		{
			name: "success",
			fields: fields{
				usecase: &MockUsecases2{},
			},
			requestbody: []byte(`
			
			{
				"seller_group_id" : 100,
				"assigment_date" : "2023-08-10 15:31:45",
				"start_date" : "2023-08-10 15:31:45",
				"due_date" : "2023-08-15 15:31:45",
				"count_success" : 100,
				"count_failed" : 0,
				"process_time" : 10000,
				"created_by" : 123
				
			}
				
			`),
			wantStatus: http.StatusOK,
			want:       []byte(`{"Code":200,"Status":"OK","Message":"Ok"}`),
			expect: []*gomock.Call{
				mockTime.EXPECT().Now().Return(time.Date(2023, 5, 8, 1, 1, 1, 1, time.Local)),
			},
		},
		{
			name: "broken json",
			fields: fields{
				usecase: &MockUsecases2{},
			},
			requestbody: []byte(`
			a
				
			`),
			wantStatus: http.StatusBadRequest,
			want:       []byte(`{"Code":400,"Status":"Bad Request","Message":"invalid_BodyRequest"}`),
			expect: []*gomock.Call{
				mockTime.EXPECT().Now().Return(time.Date(2023, 5, 8, 1, 1, 1, 1, time.Local)),
			},
		},
		{
			name:        "error empty param",
			fields:      fields{},
			requestbody: []byte(`{}`),
			wantStatus:  http.StatusBadRequest,
			want:        []byte(`{"Code":400,"Status":"Bad Request","Message":"empty_SellerGroupID"}`),
			expect: []*gomock.Call{
				mockTime.EXPECT().Now().Return(time.Date(2023, 5, 8, 1, 1, 1, 1, time.Local)),
			},
		},
		{
			name: "error usecase",
			fields: fields{
				usecase: &MockUsecases2{
					errInsertAssigmentHistory: errors.New("any"),
				},
			},

			requestbody: []byte(`
			
			{
				"seller_group_id" : 100,
				"assigment_date" : "2023-08-10 15:31:45",
				"start_date" : "2023-08-10 15:31:45",
				"due_date" : "2023-08-15 15:31:45",
				"count_success" : 100,
				"count_failed" : 0,
				"process_time" : 10000,
				"created_by" : 123
				
			}
				
			`),
			wantStatus: http.StatusInternalServerError,
			want:       []byte(`{"Code":500,"Status":"Internal Server Error","Message":"error_InsertAssigmentHistory"}`),
			expect: []*gomock.Call{
				mockTime.EXPECT().Now().Return(time.Date(2023, 5, 8, 1, 1, 1, 1, time.Local)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.InitTestingConfig()
			router := chi.NewRouter()
			hd := New(router, tt.fields.usecase)
			urlPath := "/login"
			hd.router.Post(urlPath, hd.Login)

			urlParsed, _ := url.Parse(urlPath)
			urlQuery := urlParsed.Query()

			urlParsed.RawQuery = urlQuery.Encode()

			req := httptest.NewRequest("POST", urlParsed.String(), bytes.NewBuffer(tt.requestbody))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			resp := w.Result()
			response, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}
			resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Error("invalid response")
			}

			if string(response) != string(tt.want) {
				t.Errorf("Negative - want resp [%+v] got [%+v]", string(tt.want), string(response))
			}
		})
	}
}
