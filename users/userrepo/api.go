package userrepo

type UserRepository interface {

	StoreUser(name string, email string)

	FindUserByEmail(email string)

}

func New(environment string) UserRepository {
	if environment == "production" {
		return cloudUserRepo{}
	}
	if environment == "local" {
		return sqliteUserRepo{}
	}

	return mockUserRepo{}
}