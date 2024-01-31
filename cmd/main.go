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

func main() {
	db, err := sql.Open("sqlite3", "./socialBuddies.db")
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
