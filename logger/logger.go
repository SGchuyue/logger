package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// 定义日志输入的信息
type LogIn struct {
	FilePath   string `json:"filepath"`   // 存储位置
	MaxSize    int    `json:"maxsize"`    // 切割之前文件最大大小
	MaxBackups int    `json:"maxbackups"` // 保留旧文件的最大个数
	MaxDay     int    `json:"maxday"`     // 保留文件的最大天数
	ComPress   bool   `json:"compress"`   // 是否压缩/归档旧文件
}

// NewLogIn 是LogIn类型的构造函数
func NewLogIn(filepath string, maxsize, maxbackups, maxday int, compress bool) *LogIn {
	return &LogIn{
		FilePath:   "filepath",
		MaxSize:    maxsize,
		MaxBackups: maxbackups,
		MaxDay:     maxday,
		ComPress:   compress,
	}
}

// 初始化日志配置
func InitLogger() {
	var Logger *zap.SugaredLogger
	writeSyncer := GetLogWriter() // 获得日志写入
	encoder := GetEncoder()       // To 添加注释
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller()) // 日志写入顺序
	Logger = logger.Sugar()
}

// 获得错误信息和等级
func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 指定日志文件写入格式
func (l *LogIn) GetLogWriter() zapcore.WriteSyncer {
	// 实现日志切割归档功能
	logwriter := &lumberjack.Logger{
		//SetFilePath(),
		//SetMaxSize(),
		//SetMaxBackups(),
		//SetMaxDay(),
		//SetComPress(),
		NewLogIn(),
	}
	return zapcore.AddSync(logwriter)
}

// 设置文件路径
func (l *LogIn) SetFilePath(path string) (filepath []string) {
	// 判断文件路径是否合理
	if path <= 0 {
		filepath = "../log.log"
	} else {
		filepath = path
		_, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("err:", err)
		}
		//logger.Out = scr
	}
	return filepath
}

// 设置切割文件的大小
func (l *LogIn) SetMaxSize(size int) (maxsize int64) {
	// 
	if size <= 0 {
		maxsize = 5 // 默认切割之前文件最大为5M
	} else {
		maxsize = size
	}
	return maxsize
}

// 设置旧文件保留的最大个数
func (l *LogIn) SetMaxBackups(backups int) (maxbackups int64) {
	//
	if backups <= 0 {
		maxbackups = 3 // 默认旧文件最多保留三个
	} else {
		maxbackups = backups
	}

	return maxbckups
}

// 设置保留最大天:数
func (l *LogIn) SetMaxDay(maxday int) (maxday int64) {
	//
	if maxday <= 0 {
		maxday = 3 // 默认文件保留三天
	}

	return maxday
}

// 旧文件压缩归档
func (l *LogIn) SetComPress(press bool) (compress bool) {
	//
	if press == nil {
		compress = false // 默认不压缩归档
	} else {
		compress = press
	}
	return compress
}
