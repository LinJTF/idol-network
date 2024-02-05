package user

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func (m *mockService) GetUsers(ctx context.Context) ([]User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]User), args.Error(1)
}

func (m *mockService) GetUserByID(ctx context.Context, id int) (*User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *mockService) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*User), args.Error(1)
}

func (m *mockService) CreateUser(ctx context.Context, user User) (User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(User), args.Error(1)
}

func TestHandler_GetUsers(t *testing.T) {
	mockService := &mockService{}

	tests := []struct {
		name       string
		mockResult []User
		mockErr    error
	}{
		{
			name:       "GetUsers_Success",
			mockResult: []User{{ID: 1, Name: "Lin", Email: "lin@vocedm.com.br"}},
			mockErr:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.On("GetUsers", context.Background()).Return(tt.mockResult, tt.mockErr)

			req := httptest.NewRequest("GET", "/users", nil)
			rr := httptest.NewRecorder()

			handler := NewUserHandler(mockService)
			handler.GetUsers(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
			}

			if tt.mockErr == nil {
				var gotUsers []User
				err := json.Unmarshal(rr.Body.Bytes(), &gotUsers)
				if err != nil {
					t.Fatalf("could not unmarshal response body: %v", err)
				}

				if !reflect.DeepEqual(gotUsers, tt.mockResult) {
					t.Errorf("handler returned unexpected body: got %v, want %v", gotUsers, tt.mockResult)
				}
			} else {
				if gotBody := rr.Body.String(); gotBody != tt.mockErr.Error() {
					t.Errorf("handler returned unexpected error: got %v, want %v", gotBody, tt.mockErr)
				}
			}
		})
	}
}
