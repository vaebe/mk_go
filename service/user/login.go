package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"mk/global"
	middlewares "mk/middleware"
	"mk/models"
	"mk/service/email"
	"mk/utils"
	"time"
)

// generateToken 生成token
func generateToken(user models.User) (token string, err error) {
	//生成token
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint(user.ID),
		NickName:    user.NickName,
		AuthorityId: uint(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 签名的生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24小时
			Issuer:    "1057",
		},
	}
	return j.CreateToken(claims)
}

// SendVerificationCode
// @Summary     发送验证码
// @Description  发送验证码
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 			param body    models.VerificationCodeForm  true  "请求对象"
// @Success      200  {object}  utils.ResponseResultInfo
// @Failure      500  {object}  utils.EmptyInfo
// @Router       /user/sendVerificationCode [post]
func SendVerificationCode(ctx *gin.Context) {
	//表单验证
	verificationCodeForm := models.VerificationCodeForm{}

	if err := ctx.ShouldBind(&verificationCodeForm); err != nil {
		zap.S().Info(&verificationCodeForm)
		utils.HandleValidatorError(ctx, err)
		return
	}

	// 获取随机验证码
	verificationCode := utils.GenerateSmsCode(6)

	// 发送验证码邮件
	err := email.SendTheVerificationCodeEmail(verificationCode, verificationCodeForm.Email)
	if err != nil {
		utils.ResponseResultsError(ctx, "发送邮件验证码失败")
	}

	// 将数据存储到redis
	global.RedisClient.Set(context.Background(), verificationCodeForm.Email, verificationCode, time.Duration(global.RedisConfig.Expire)*time.Second)
	utils.ResponseResultsSuccess(ctx, "发送验证码成功！")
}

// loginSuccess 登陆成功后的操作
func loginSuccess(ctx *gin.Context, user models.User) {
	token, err := generateToken(user)
	if err != nil {
		zap.S().Info("生成token错误", err.Error())
		utils.ResponseResultsError(ctx, "生成token错误!")
		return
	}

	resultsData := map[string]any{
		"id":         user.ID,
		"nick_name":  user.NickName,
		"github":     user.Github,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	}
	utils.ResponseResultsSuccess(ctx, resultsData)
}

// Register
// @Summary     用户注册
// @Description  用户注册
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 			param body    models.RegisterForm  true  "请求对象"
// @Success      200  {object}  utils.ResponseResultInfo
// @Failure      500  {object}  utils.EmptyInfo
// @Router       /user/register [post]
func Register(ctx *gin.Context) {
	//表单验证
	registerForm := models.RegisterForm{}

	if err := ctx.ShouldBind(&registerForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	// 验证短信验证码
	redisCode := global.RedisClient.Get(context.Background(), registerForm.Email)
	verificationCode, _ := redisCode.Result()

	if verificationCode == "redis" || verificationCode != registerForm.Code {
		zap.S().Infof("验证码不正确:应为%s实际为%s", verificationCode, registerForm.Code)

		utils.ResponseResultsError(ctx, "验证码不正确!")
		return
	}

	user := models.User{
		NickName:    fmt.Sprintf("mk%s", uuid.New().String()),
		UserAvatar:  "https://foruda.gitee.com/avatar/1677018140565464033/3040380_mucuni_1578973546.png",
		UserAccount: registerForm.Email, // 暂时使用邮箱注册
		Password:    registerForm.PassWord,
	}
	userRes := global.DB.Create(&user)

	if userRes.Error != nil {
		zap.S().Info("创建用户失败！", userRes.Error)
		utils.ResponseResultsError(ctx, userRes.Error.Error())
		return
	}

	loginSuccess(ctx, user)
}

// Login
// @Summary     用户登陆
// @Description  用户登陆
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 			param body    models.LoginForm  true  "请求对象"
// @Success      200  {object}  utils.ResponseResultInfo
// @Failure      500  {object}  utils.EmptyInfo
// @Router       /user/login [post]
func Login(ctx *gin.Context) {
	//表单验证
	loginForm := models.LoginForm{}

	if err := ctx.ShouldBind(&loginForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	userInfo := models.User{}

	global.DB.Model(&models.User{UserAccount: loginForm.Email}).First(&userInfo)

	if userInfo.Password != loginForm.PassWord {
		utils.ResponseResultsError(ctx, "密码不正确")
		return
	}

	loginSuccess(ctx, userInfo)
}
