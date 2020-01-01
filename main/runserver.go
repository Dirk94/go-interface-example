package main

import "github.com/dirk94/go-medium.git/users/userrepo"

func main() {
	// In the real world you would probably get this as a
	// command line argument, or read it from a dotenv file.
	environment := "production"

	repo := userrepo.New(environment)

	repo.StoreUser("Dirk", "dirk@some-email.com")
}