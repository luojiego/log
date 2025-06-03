package main

import (
	log "github.com/luojiego/slogx"
)

func logSomething() {
	log.Debug("This is a debug message")
	log.Info("This is an info message")

	// 测试带字段的日志
	log.Info("User logged in", "userId", 123, "ip", "192.168.1.1")
}

func main() {
	// 直接调用包级别的函数
	log.Info("Application started")

	// 在不同的函数中调用
	logSomething()

	// 测试 With 功能
	logger := log.With("module", "auth")
	logger.Error("Authentication failed", "reason", "invalid_token")

	// 测试不同级别的日志
	log.Debug("Debug level message")
	log.Info("Info level message")
	log.Warn("Warning level message")
	log.Error("Error level message")
}
