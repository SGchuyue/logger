package main

import (
	"logger/logger"
)

func main() {
	// 对日志包进行初始化
	logger.InitLogger("test.log", 2, 3, 5, true)
	defer logger.sugarLogger.Sync()
	logger.Error("test error %s", "test")
	//logger(fmt.Println(time.Now()))
	//logger("hello")

}
