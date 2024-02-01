package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"socialbuddies/internal/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func startSqliteDb(tableName string, databasePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Name TEXT,
			Email TEXT,
			BirthDate TEXT,
			Phone TEXT,
			DocumentNumber TEXT,
			Street TEXT,
			Number TEXT,
			Complement TEXT,
			City TEXT,
			Country TEXT,
			State TEXT,
			ZipCode TEXT
		)
	`, tableName))
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := startSqliteDb("Users", "./internal/db/socialBuddies.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := user.NewUserRepository(db)
	userService := user.NewService(userRepository)
	userHandler := user.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/v1/user", userHandler.GetUsers)
	r.Get("/v1/user/{id}", userHandler.GetUserByID)
	r.Get("/v1/user/email/{email}", userHandler.GetUserByEmail)
	r.Post("/v1/user", userHandler.CreateUser)

	fmt.Println("Server listening on port 8081...")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal(err)
	}
}

//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Println("Moses")
//}
