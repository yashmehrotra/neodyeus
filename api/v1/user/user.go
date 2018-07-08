package user

type User struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

func CreateUser(u User) error {
}

func AuthUser(email, password string) bool {
}
