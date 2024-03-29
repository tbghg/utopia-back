package v1

import (
	"errors"
	"github.com/google/uuid"
	"os"
	"path"
	"utopia-back/database/abstract"
	"utopia-back/model"
	utils "utopia-back/pkg/util"
)

type UserService struct {
	UserDal abstract.UserDal
}

func (u *UserService) GetUserInfo(id uint) (userInfo model.UserInfo, err error) {
	return u.UserDal.GetUserInfoById(id)
}

var (
	ErrorUserExists    = errors.New("用户已存在")
	ErrorUserNotExists = errors.New("用户不存在")
	ErrorPasswordWrong = errors.New("密码错误")
)

func (u *UserService) Login(username string, password string) (token string, id uint, err error) {
	// 检查用户是否存在
	user, exist, err := u.UserDal.GetUserByUsername(username)
	if !exist {
		return "", 0, ErrorUserNotExists
	}
	if err != nil {
		return "", 0, err
	}
	// 检查密码是否正确
	if utils.ValidMd5EncodeWithSalt(password, user.Salt, user.Password) {
		//  生成token
		return utils.GenToken(int(user.ID)), user.ID, nil
	}
	return "", 0, ErrorPasswordWrong
}

func (u *UserService) Register(username string, password string) (token string, id uint, err error) {
	// 检查用户是否存在
	_, exist, err := u.UserDal.GetUserByUsername(username)
	if exist {
		return "", 0, ErrorUserExists
	}
	if err != nil {
		return "", 0, err
	}

	// 密码加密+盐
	salt := utils.RandomSalt()
	password = utils.Md5EncodeWithSalt(password, salt)
	user := model.User{Username: username, Password: password, Salt: salt}
	// 生成随机昵称
	uuidStr := uuid.New().String()[:8]
	user.Nickname = uuidStr

	//生成头像
	avatarName := uuidStr + ".png"                      // 头像名称
	nativeAvatarPath := path.Join("output", avatarName) // 头像保存路径
	err = utils.QuickGenAvatar(nativeAvatarPath)        // 生成头像
	if err != nil {
		return "", 0, err
	}
	newAvatarPath, err := utils.QuickUploadFile(nativeAvatarPath, avatarName) // 上传头像到七牛云kodo对象存储
	if err != nil {
		return "", 0, err
	}
	user.Avatar = newAvatarPath       // 保存头像访问路径
	err = os.Remove(nativeAvatarPath) // 删除本地头像

	// 创建用户
	id, err = u.UserDal.CreateUser(&user)
	if err != nil {
		return "", 0, err
	}
	// 生成token
	return utils.GenToken(int(id)), user.ID, nil

}

func (u *UserService) UpdateNickname(uid uint, nickname string) (err error) {
	return u.UserDal.UpdateNickname(uid, nickname)
}
