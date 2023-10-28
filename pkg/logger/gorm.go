package logger

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"runtime"
	"strings"
	"time"
)

// 手动实现gorm的logger接口
var _ gormLogger.Interface = (*GormLogger)(nil)

type GormLogger struct {
	zapLogger  *zap.Logger
	ShowThread time.Duration
	logLevel   gormLogger.LogLevel
}

func NewGormLogger(Logger *zap.Logger, level gormLogger.LogLevel) *GormLogger {

	return &GormLogger{
		logLevel:   level,
		zapLogger:  Logger,
		ShowThread: 200 * time.Millisecond, //慢查询阈值 默认200ms
	}
}

func (g GormLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return GormLogger{
		zapLogger:  g.zapLogger,
		ShowThread: g.ShowThread,
		logLevel:   level,
	}
}

func (g GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	if g.logLevel < gormLogger.Info {
		return
	}

	g.logger().Sugar().With(ctx).Infof(s, i...)
}

func (g GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if g.logLevel < gormLogger.Warn {
		return
	}
	g.logger().Sugar().With(ctx).Warnf(s, i...)
}

func (g GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if g.logLevel < gormLogger.Error {
		return
	}
	g.logger().Sugar().With(ctx).Errorf(s, i...)
}

func (g GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//获取运行时间
	elapsed := time.Since(begin)
	//获取sql
	sql, rows := fc()
	//字段
	fields := []zap.Field{
		zap.String("sql", sql),
		zap.String("elapsed", elapsed.String()),
		zap.Int64("rows", rows),
	}
	//判断是否有错误
	if err != nil {
		//判断是否是记录不存在的错误 如果是记录不存在的错误 Warn级别
		if errors.Is(err, gorm.ErrRecordNotFound) {
			g.zapLogger.Warn("record not found", fields...)
			return
		} else {
			//其他错误 Error级别
			logFields := append(fields, zap.Error(err))
			g.zapLogger.Error("gorm error", logFields...)
		}
	}

	//判断是否是慢查询
	if g.ShowThread != 0 && elapsed > g.ShowThread {
		g.logger().Warn("slow query", fields...)
	}

	//只有在Info级别才打印所有的sql
	if g.logLevel >= gormLogger.Info {
		g.logger().Debug("every query", fields...)
	}

}

func (g GormLogger) logger() *zap.Logger {
	//跳过gorm的调用栈
	var (
		gormPackage    = "gorm.io/gorm"
		zapGormPackage = "moul.io/zapgorm2"
	)
	//减少一层调用栈及初始化中的zap.AddCallerSkip(1)
	clone := g.zapLogger.WithOptions(zap.AddCallerSkip(-2))
	//获取调用栈
	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapGormPackage):
		default:
			return clone.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return g.zapLogger

}
