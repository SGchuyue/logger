package logger_test

import (
	"logger/logger"
	"testing"
)

// 测试logger日志包的功能
func TestInitLogger(t *testing.T) {
	logger.InitLogger("test.log", 2, 3, 5, true)
	for i := 0; i < 50000; i++ {
		logger.Error("test  %s", "test")
	}
	// PASS
	//ok      logger/logger   0.593s
}
