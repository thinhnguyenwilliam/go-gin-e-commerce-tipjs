package repo

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUserRepo() User {
	return User{
		Name:  "Mr William Handsome",
		Email: "william@example.com",
		Age:   30,
	}
}
