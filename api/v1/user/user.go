package user

type User struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

func createUser(u User) error {
	// Check if email exists in DB
	// If not, return error
	// else create user
	return nil
}

func AuthUser(email, password string) bool {
	return true
}
