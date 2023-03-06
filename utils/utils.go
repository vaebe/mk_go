package utils

import (
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"net"
	"strings"
	"time"
)

// GetFreePort 获取可用端口
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer func(l *net.TCPListener) {
		err := l.Close()
		if err != nil {
			zap.S().Error("获取可用端口信息失败！")
		}
	}(l)
	return l.Addr().(*net.TCPAddr).Port, nil
}

// GenerateSmsCode 生成width长度的短信验证码
func GenerateSmsCode(width int) string {

	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.NewSource(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			return ""
		}
	}
	return sb.String()
}
