package user

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type UserRepository interface {
	GetUsers() ([]User, error)
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user User) (User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUsers() ([]User, error) {
	rows, err := r.db.Query("SELECT * FROM Users")
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

func (r *userRepository) GetUserByID(id int) (*User, error) {
	var user User
	err := r.db.QueryRow("SELECT * FROM Users WHERE ID = ?", id).Scan(
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
	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.db.QueryRow("SELECT * FROM Users WHERE Email = ?", email).Scan(
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

	return &user, nil
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result, err := r.db.Exec(`
        INSERT INTO Users (
            Name,
            Email,
            BirthDate,
            Phone,
            DocumentNumber,
            Street,
            Number,
            Complement,
            City,
            Country,
            State,
            ZipCode
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `,
		user.Name,
		user.Email,
		user.BirthDate,
		user.Phone,
		user.DocumentNumber,
		user.Address.Street,
		user.Address.Number,
		user.Address.Complement,
		user.Address.City,
		user.Address.Country,
		user.Address.State,
		user.Address.ZipCode,
	)

	if err != nil {
		return User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}

	var newUser User
	err = r.db.QueryRow("SELECT * FROM Users WHERE ID = ?", id).Scan(
		&newUser.ID,
		&newUser.Name,
		&newUser.Email,
		&newUser.BirthDate,
		&newUser.Phone,
		&newUser.DocumentNumber,
		&newUser.Address.Street,
		&newUser.Address.Number,
		&newUser.Address.Complement,
		&newUser.Address.City,
		&newUser.Address.Country,
		&newUser.Address.State,
		&newUser.Address.ZipCode,
	)
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}
