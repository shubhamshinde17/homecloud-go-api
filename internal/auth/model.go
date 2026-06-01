package auth

type User struct {
	ID           int64
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	IsActive     bool
}
