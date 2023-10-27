package abstract

type UserService interface {
	// Login 登录
	Login(username string, password string) (token string, id uint, err error)
	// Register 注册
	Register(username string, password string) (token string, id uint, err error)
}
