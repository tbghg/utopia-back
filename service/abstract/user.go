package abstract

import "utopia-back/model"

type UserService interface {
	// Login 登录
	Login(username string, password string) (token string, id uint, err error)
	// Register 注册
	Register(username string, password string) (token string, id uint, err error)
	// GetUserInfo 获取用户信息
	GetUserInfo(id uint) (userInfo model.UserInfo, err error)
	// UpdateNickname 修改昵称
	UpdateNickname(uid uint, nickname string) error
}
