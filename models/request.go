package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	ID          uint
	UserId      uint
	NickName    string
	AuthorityId uint
	jwt.RegisteredClaims
}
