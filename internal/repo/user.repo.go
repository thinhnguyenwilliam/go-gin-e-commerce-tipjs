package repo

type IUserRepo interface {
	GetUserByEmail(email string) bool
}

type userRepo struct{}

// Constructor to create a new UserRepo instance
func NewUserRepo() IUserRepo {
	return &userRepo{}
}

// Dummy implementation — replace with actual DB query
func (r *userRepo) GetUserByEmail(email string) bool {
	// TODO: query database here
	return false
}

// type User struct {
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// 	Age   int    `json:"age"`
// }

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUserRepo() User {
// 	return User{
// 		Name:  "Mr William Handsome",
// 		Email: "william@example.com",
// 		Age:   30,
// 	}
// }
