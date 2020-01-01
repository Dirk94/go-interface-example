package userrepo

import "fmt"

type cloudUserRepo struct {
}

func (m cloudUserRepo) StoreUser(name string, email string) {
	fmt.Println("cloudUserRepo: mocking the StoreUser func")
	// In the real world you would store the user in the Google Cloud database here.
}

func (m cloudUserRepo) FindUserByEmail(email string) {
	fmt.Println("cloudUserRepo: mocking the FindUserByEmail func")
	// In the real world you would query the Google Cloud database here.
}
