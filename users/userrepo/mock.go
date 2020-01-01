package userrepo

import "fmt"

type mockUserRepo struct {
}

func (m mockUserRepo) StoreUser(name string, email string) {
	fmt.Println("mocking the StoreUser func")
}

func (m mockUserRepo) FindUserByEmail(email string) {
	fmt.Println("mocking the FindUserByEmail func")
}
