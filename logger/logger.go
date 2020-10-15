package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

// 初始化日志配置
func InitLogger(filename string, maxsize, maxbackups, maxday int, compress bool) {
	// filename 设置存入日志文件路径及名称 字符串类型 （示例：./test.log）
	// maxsize 配置文件大小 int类型 （示例：2M）
	// maxbackups 配置旧文件保留个数 int类型 （示例： 3个）
	// maxday 配置日志文件保留最大天数 int类型 （示例： 5天）
	// compress 配置文件是否进行压缩 bool类型 （fasle不压缩，ture压缩）
	getLogWriter := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxsize,
		MaxBackups: maxbackups,
		MaxAge:     maxday,
		Compress:   compress,
	}

	writeSyncer := zapcore.AddSync(getLogWriter) // 获得日志写入
	encoder := getEncoder()                      // To 添加注释
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller()) // 日志写入顺序
	sugarLogger = logger.Sugar()
}

// 获得错误信息和等级
func getEncoder() zapcore.Encoder {
	// 配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func Error(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args)
}