package userRep

type AuthRepository interface {
	Log(email string, password string) ()
	Reg(username string, email string, password string)

	// HashPassword(u *entity.User) (string, error)
	// GenerateToken(userId string) (string, error)
	// ComparePassword(password string, u *entity.User) bool
	// SendTokenResponse(u *entity.User) (string, error)
}
