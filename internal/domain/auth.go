package domain

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type AuthUsecase interface {
	Login(username, password string) (string, error)
	Register(username, password string) error
}

type AuthRepository interface {
	GetByUsername(username string) (*User, error)
	CreateUser(username, hashedPassword string) error
}
