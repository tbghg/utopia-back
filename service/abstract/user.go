package abstract

type UserService interface {
	Login(username string, password string) (token string, id uint, err error)
	Register(username string, password string) (token string, id uint, err error)
}
