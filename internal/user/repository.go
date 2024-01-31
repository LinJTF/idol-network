package user

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type UserRepository interface {
	GetUsers() ([]User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUsers() ([]User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.BirthDate,
			&user.Phone,
			&user.DocumentNumber,
			&user.Address.Street,
			&user.Address.Number,
			&user.Address.Complement,
			&user.Address.City,
			&user.Address.Country,
			&user.Address.State,
			&user.Address.ZipCode,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
