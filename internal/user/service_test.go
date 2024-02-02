package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) GetUsers() ([]User, error) {
	args := m.Called()
	return args.Get(0).([]User), args.Error(1)
}

func (m *mockUserRepository) GetUserByID(id int) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *mockUserRepository) GetUserByEmail(email string) (*User, error) {
	args := m.Called(email)
	return args.Get(0).(*User), args.Error(1)
}

func (m *mockUserRepository) CreateUser(user User) (User, error) {
	args := m.Called(user)
	return args.Get(0).(User), args.Error(1)
}

func TestService_GetUsers(t *testing.T) {
	mockRepo := &mockUserRepository{}
	svc := NewService(mockRepo)

	mockUsers := []User{
		{ID: 1, Name: "Xuxu", Email: "xuxu@vocedm.com.br"},
		{ID: 2, Name: "Lin", Email: "lin@vocedm.com.br"},
	}

	mockRepo.On("GetUsers").Return(mockUsers, nil)

	tests := []struct {
		name      string
		wantUsers []User
		wantErr   error
	}{
		{
			name:      "Suceess when getting users from repository",
			wantUsers: mockUsers,
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			users, err := svc.GetUsers(context.Background())

			if !reflect.DeepEqual(users, tt.wantUsers) {
				t.Errorf("got users %+v, want %+v", users, tt.wantUsers)
			}

			if (err != nil && tt.wantErr == nil) || (err == nil && tt.wantErr != nil) {
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
			}
		})
	}
}
