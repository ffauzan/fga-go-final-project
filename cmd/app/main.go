package main

import (
	"final-project/pkg/auth"
	"final-project/pkg/crypto"
	"final-project/pkg/http/rest"
	"final-project/pkg/storage/sqldb"
	"final-project/pkg/user"
	"net/http"
	"os"

	"log"
)

func main() {
	PORT := "8080"
	// This sensitive information is written here for the convenience of this assignment
	dsn := "falfal:Pasword!2@tcp(mysql-dev-db.airy.my.id:3306)/fga_go_final?charset=utf8mb4&parseTime=True&loc=Local"
	os.Setenv("JWT_SECRET", "supersecret1287401bnf9147ehfn9r247")

	// Create storage
	storage, err := sqldb.NewStorage(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	// Create user repository
	userRepo := sqldb.NewUserRepository(storage.DB)

	// Create Auth service
	authService := auth.NewAuthService()

	// Create Crypto service
	cryptoService := crypto.NewCryptoService()

	// Create user service
	userService := user.NewService(userRepo, cryptoService, authService)

	// Create router
	router := rest.NewRouter(
		&userService,
		&authService,
	)

	// Start server
	log.Println("Starting server on port " + PORT)
	http.ListenAndServe(":"+PORT, router)

}
