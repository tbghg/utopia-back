package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"time"
)

var Logger *zap.Logger

func InitLogger(maxSize, maxBackup, maxAge int, compress bool, level string) {

	//设置写入介质
	writeSyncer := getLogWriter(maxSize, maxBackup, maxAge, compress)
	//设置日志级别

	var logLevel zapcore.Level
	switch level {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	default:
		logLevel = zap.InfoLevel
	}
	//初始化core
	core := zapcore.NewCore(zapcore.NewJSONEncoder(getEncoderConfig()), writeSyncer, logLevel)

	Logger = zap.New(core,
		zap.AddCaller(),                   //调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),              //封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel)) //Error 时才会显示 stacktrace

}

// 获取日志格式
func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",                         // 时间格式
		LevelKey:       "level",                        // 日志级别
		NameKey:        "logger",                       // 日志名称
		CallerKey:      "caller",                       // 代码调用，如 paginator/paginator.go:148
		FunctionKey:    zapcore.OmitKey,                // 函数名称
		MessageKey:     "message",                      // 日志信息
		StacktraceKey:  "stacktrace",                   // 堆栈信息
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写,并且添加颜色
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, // 时间格式
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 调用文件和行号，内部使用 runtime.Caller
	}
}

const loggerPath = "log"

// 获取日志写入介质
func getLogWriter(size int, backup int, age int, compress bool) zapcore.WriteSyncer {
	logName := path.Join(loggerPath, fmt.Sprintf(time.Now().Format("2006-01-02")+".log"))
	//使用lumberjack进行日志分割归档
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logName,
		MaxSize:    size,
		MaxBackups: backup,
		MaxAge:     age,
		Compress:   compress,
	}
	//写入文件和控制台
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))

}
