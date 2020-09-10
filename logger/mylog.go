package logger

import (
        "net/http"

        "github.com/natefinch/lumberjack"
        "go.uber.org/zap"
        "go.uber.org/zap/zapcore"
)

var sLogger *zap.SugaredLogger

// 初始化日志配置
func InitLogger() {
        writeSyncer := GetLogWriter() // 获得日志写入
        encoder := GetEncoder() // 
        core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

        logger := zap.New(core, zap.AddCaller()) // 日志写入顺序
        sLogger = logger.Sugar()
}

// 
func GetEncoder() zapcore.Encoder {
        encoderConfig := zap.NewProductionEncoderConfig()
        encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
        encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
        return zapcore.NewConsoleEncoder(encoderConfig)
}

// 指定日志文件写入格式
func GetLogWriter() zapcore.WriteSyncer {
        // 借用第三方库实现日志切割归档功能
        lumberJackLogger := &lumberjack.Logger{
                Filename:   "./log.log", // 存储位置
                MaxSize:    1,          // 切割之前文件最大大小
                MaxBackups: 5,          // 保留旧文件的最大个数
                MaxAge:     30,         // 保留文件的最大天数
                Compress:   false,      // 是否压缩/归档旧文件
        }
        return zapcore.AddSync(lumberJackLogger)
}

// 获取网页信息
func SimpleHttpGet(url string) {
        sLogger.Debugf("正在尝试获得 %s 的请求", url)
        resp, err := http.Get(url)
        if err != nil {
                // 打印错误信息
                sLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
                sLogger.Errorf("获取URL%s时出错，错误信息为%s", url, err)
        } else {
                // 
                sLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
                sLogger.Infof("成功获取！状态码=%s URL = %s", resp.Status, url)
                resp.Body.Close()
        }
}
