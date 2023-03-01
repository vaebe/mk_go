package utils

import (
	"go.uber.org/zap"
	"net"
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
