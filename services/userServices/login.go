package userServices

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"mk/global"
	"mk/models/user"
)

// Register 用户注册
func Register(registerForm user.RegisterForm) (user.User, error) {
	// 验证短信验证码
	redisCode := global.RedisClient.Get(context.Background(), registerForm.Email)
	verificationCode, _ := redisCode.Result()

	if verificationCode == "redis" || verificationCode != registerForm.Code {
		zap.S().Infof("验证码不正确:应为%s实际为%s", verificationCode, registerForm.Code)

		return user.User{}, errors.New("验证码不正确")
	}

	// 生成不带 - 的uuid
	uuidObj := uuid.New()
	uuidStr := fmt.Sprintf("mk%x", uuidObj[:])

	userInfo := user.User{
		NickName:    uuidStr,
		UserName:    uuidStr,
		UserAvatar:  "https://foruda.gitee.com/avatar/1677018140565464033/3040380_mucuni_1578973546.png",
		UserAccount: registerForm.Email, // 暂时使用邮箱注册
		Password:    registerForm.PassWord,
	}
	res := global.DB.Create(&userInfo)

	if res.Error != nil {
		return user.User{}, res.Error
	}

	return userInfo, nil
}

// VerifyUserPassword 校验用户密码
func VerifyUserPassword(loginForm user.LoginForm) (user.User, error) {
	userInfo := user.User{}

	res := global.DB.Where("user_account = ?", loginForm.Email).First(&userInfo)

	if res.Error != nil {
		return user.User{}, res.Error
	}

	if userInfo.Password != loginForm.PassWord {
		return user.User{}, errors.New("密码不正确")
	}

	userInfo.Password = ""

	return userInfo, nil
}
