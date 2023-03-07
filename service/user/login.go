package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"mk/global"
	middlewares "mk/middleware"
	"mk/models"
	"mk/service/email"
	"mk/utils"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var users []models.UserTest
	res := global.DB.Find(&users)

	// 存在错误
	if res.Error != nil {
		return
	}

	zap.S().Info("数据", users)

	data := map[string]interface{}{
		"code": "0",
		"msg":  "登陆",
	}

	// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	c.AsciiJSON(http.StatusOK, data)
}

// SendVerificationCode 发送验证码
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

// Register
// @Summary     用户注册
// @Description  用户注册
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 			param body    models.RegisterForm  true  "请求对象"
// @Success      200  {object}  utils.ResponseResultInfo
// @Failure      400  {object}  utils.EmptyInfo
// @Failure      404  {object}  utils.EmptyInfo
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
	redisCode := global.RedisClient.Get(context.Background(), registerForm.Mobile)
	verificationCode, _ := redisCode.Result()

	if verificationCode == "redis" || verificationCode != registerForm.Code {
		zap.S().Info("验证码不正确！")

		utils.ResponseResultsError(ctx, "验证码不正确!")
		return
	}

	user := models.UserTest{
		NickName: registerForm.Mobile,
		Password: registerForm.PassWord,
		Mobile:   registerForm.Mobile,
		Role:     int(1),
	}
	userRes := global.DB.Create(&user)

	if userRes.Error != nil {
		zap.S().Info("创建用户失败！", userRes.Error)
		utils.ResponseResultsError(ctx, userRes.Error.Error())
		return
	}

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
	token, err := j.CreateToken(claims)

	if err != nil {
		zap.S().Info("生成token错误", err.Error())
		utils.ResponseResultsError(ctx, "生成token错误!")
		return
	}

	resultsData := map[string]any{
		"id":         user.ID,
		"nick_name":  user.NickName,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	}
	utils.ResponseResultsSuccess(ctx, resultsData)
}
