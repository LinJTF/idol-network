package user

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	ErrNoRows = errors.New("no rows in result set")
)

func TestRepository_GetUsers(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	repo := NewUserRepository(mockDB)

	mockRows := sqlmock.NewRows([]string{"ID", "Name", "Email", "BirthDate", "Phone", "DocumentNumber", "Street", "Number", "Complement", "City", "Country", "State", "ZipCode"}).
		AddRow(1, "Xuxu", "xuxu@vocedm.com.br", "2000-01-02", "123456789", "123456789", "Rua dos brabos", "123", "Apt 1", "Abacaxi", "Fenda do Bikini", "Pernambuco", "12345").
		AddRow(2, "Lin", "lin@vocedm.com.br", "2000-10-17", "88888888", "12345678987654", "Rua do Chines", "123", "Apt 1", "Hellcife", "Uga Buga", "Antidoggos", "12345")

	mock.ExpectQuery("SELECT \\* FROM Users").WillReturnRows(mockRows)

	tests := []struct {
		name      string
		wantUsers []User
		wantErr   error
	}{
		{
			name: "Success when getting users from database",
			wantUsers: []User{
				{ID: 1, Name: "Xuxu", Email: "xuxu@vocedm.com.br", BirthDate: "2000-01-02", Phone: "123456789", DocumentNumber: "123456789", Address: Address{Street: "Rua dos brabos", Number: "123", Complement: "Apt 1", City: "Abacaxi", Country: "Fenda do Bikini", State: "Pernambuco", ZipCode: "12345"}},
				{ID: 2, Name: "Lin", Email: "lin@vocedm.com.br", BirthDate: "2000-10-17", Phone: "88888888", DocumentNumber: "12345678987654", Address: Address{Street: "Rua do Chines", Number: "123", Complement: "Apt 1", City: "Hellcife", Country: "Uga Buga", State: "Antidoggos", ZipCode: "12345"}},
			},
			wantErr: nil,
		},
		{
			name:      "Error when no rows in result set",
			wantUsers: nil,
			wantErr:   ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users, err := repo.GetUsers()

			if !reflect.DeepEqual(users, tt.wantUsers) {
				t.Errorf("got users %+v, want %+v", users, tt.wantUsers)
			}

			if (err != nil && tt.wantErr == nil) || (err == nil && tt.wantErr != nil) {
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
			}
		})
	}
}
