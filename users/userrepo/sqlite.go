package userrepo

import "fmt"

type sqliteUserRepo struct {
}

func (m sqliteUserRepo) StoreUser(name string, email string) {
	fmt.Println("sqliteUserRepo: mocking the StoreUser func")
	// In a real world project you would query an sqlite database here.
}

func (m sqliteUserRepo) FindUserByEmail(email string) {
	fmt.Println("sqliteUserRepo: mocking the FindUserByEmail func")
	// In a real world project you would query an sqlite database here.
}
