package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 定义*zap.SugaredLogger类型的包内变量
// 使用SugaredLogger类型的zap日志记录器
var sugarLogger *zap.SugaredLogger

// 初始化日志配置
func InitLogger(filename string, maxsize, maxbackups, maxday int, compress bool) {
	// filename 设置存入日志文件路径及名称 字符串类型
	// maxsize 配置文件大小 int类型 （单位为M）
	// maxbackups 配置旧文件保留个数 int类型
	// maxday 配置日志文件保留最大天数 int类型
	// compress 配置文件是否进行压缩 bool类型 （fasle不压缩，ture压缩）
	getLogWriter := &lumberjack.Logger{
		Filename:   filename,   // 文件路径
		MaxSize:    maxsize,    // 日志大小（M）
		MaxBackups: maxbackups, // 保留个数
		MaxAge:     maxday,     // 留存天数
		Compress:   compress,   // 是否压缩
	}
	// 指定日志将写到哪里
	//使用zapcore.AddSync()函数并且将打开的文件句柄传进去。
	writeSyncer := zapcore.AddSync(getLogWriter)
	// 得到编码和调用者信息
	encoder := getEncoder()
	// 使用zap.New(…)方法来手动传递所有配置
	// zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// 将调用函数信息记录到日志中的功能
	logger := zap.New(core, zap.AddCaller())
	// 调用主logger中的Sugar()方法来获取一个SugeredLogger
	sugarLogger = logger.Sugar()
}

// 更改时间编码并添加调用者详细信息
func getEncoder() zapcore.Encoder {
	// 覆盖默认的ProductionConfig()配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 更改时间编码配置，以人类可读的方式展示
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 获取日志等级的信息
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 将JSON Encoder更改为普通的Log Encoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 返回错误信息函数
func Error(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args)
}
