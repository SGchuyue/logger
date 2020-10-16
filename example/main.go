package main

import (
	"github.com/SGchuyue/logger/logger"
)

func main() {
	// 对日志包进行初始化
	logger.InitLogger("test.log", 2, 3, 5, true)
	// 对日志包的使用
	logger.Error("error test")
	logger.Errorf("test errorf %s", "test")
	logger.Debug("debug test")
	logger.Debugf("test debugf %s", "test")
	logger.Info("info test")
	logger.Infof("test infof %s", "test")
	logger.Warn("warn test")
	logger.Warnf("test warnf %s", "test")
	// 慎用，程序中断后面无法正常打印
	logger.Fatal("fatal test")
	// logger.Fatalf("test fatal %s","test")
	// logger.Panic("panic test")
	// logger.Panicf("test panicf %s","test")
}
