package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 使用SugaredLogger类型的zap日志记录器
var sugarLogger *zap.SugaredLogger

// InitLogger 初始化日志库的使用
// filename 设置存入日志文件路径及名称 字符串类型
// maxsize 配置文件大小 int类型 （单位为M）
// maxbackups 配置旧文件保留个数 int类型
// maxday 配置日志文件保留最大天数 int类型
// compress 配置文件是否进行压缩 bool类型 （fasle不压缩，ture压缩）
func InitLogger(filename string, maxsize, maxbackups, maxday int, compress bool) {
	getLogWriter := &lumberjack.Logger{
		Filename:   filename,   // 文件路径
		MaxSize:    maxsize,    // 日志大小（M）
		MaxBackups: maxbackups, // 保留个数
		MaxAge:     maxday,     // 留存天数
		Compress:   compress,   // 是否压缩
	}
	// 使用zapcore.AddSync()函数并且将打开的文件句柄传进去。
	writeSyncer := zapcore.AddSync(getLogWriter)
	//defer sugarLogger.Sync()
	// 得到编码和调用者信息
	encoder := getEncoder()
	// 使用zap.New(…)方法来手动传递所有配置
	// zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	// 调用主logger中的Sugar()方法来获取一个SugeredLogger
	sugarLogger = logger.Sugar()
}

// 更改时间编码并添加调用者详细信息
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 更改时间编码配置，以人类可读的方式展示
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 将JSON Encoder更改为普通的Log Encoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// Error 发生错误不影响程序运行打印错误和异常信息
func Error(args ...interface{}) {
	sugarLogger.Error(args)
}

// Errorf 使用 fmt.Sprintf 记录模板化信息
func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args)
}

// Info 突出强调应用程序的运行过程打印重要信息
func Info(template string, args ...interface{}) {
	sugarLogger.Info(template, args)
}

// Infof 使用 fmt.Sprintf 记录模板化信息
func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args)
}

// Debug 开发过程中打印一些运行信息
func Debug(args ...interface{}) {
	sugarLogger.Debug(args)
}

// Debugf 使用 fmt.Sprintf 记录模板化信息
func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args)
}

// ToDo 注释
func Panic(args ...interface{}) {
	sugarLogger.Panic(args)
}

// Panicf 使用 fmt.Sprintf 记录模板化信息
func Panicf(template string, args ...interface{}) {
	sugarLogger.Panicf(template, args)
}

// Warn 出现潜在错误的提示
func Warn(args ...interface{}) {
	sugarLogger.Warn(args)
}

// Warnf 使用 fmt.Sprintf 记录模板化信息
func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args)
}

// Fatal 会导致应用程序的退出(慎用)
func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args)
}

// Fatalf 使用 fmt.Sprintf 记录模板化信息
func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args)
}