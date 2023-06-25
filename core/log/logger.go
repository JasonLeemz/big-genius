package log

import (
	"big-genius/core/config"
	"context"
	"errors"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

// Logger 声明日志类全局变量
var Logger *zap.SugaredLogger

type GormLogger struct {
	logger *zap.SugaredLogger

	LogLevel                  logger.LogLevel
	SlowThreshold             time.Duration
	Colorful                  bool
	IgnoreRecordNotFoundError bool
	ParameterizedQueries      bool
}

var (
	infoStr      = "%s\tINFO "
	warnStr      = "%s\tWARN "
	errStr       = "%s\tERROR "
	traceStr     = "%s\t[%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s\t[%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s\t[%.3fms] [rows:%v] %s"
)

var Glogger *GormLogger

// newGormLogger GormLogger 初始化
func newGormLogger() {
	Glogger = &GormLogger{
		logger:   Logger,
		LogLevel: logger.Info,
	}

}

func (glog *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	lg := logger.Default.LogMode(level)
	return lg
}

func (glog *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if glog.LogLevel >= logger.Info {
		glog.logger.Infof(infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (glog *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if glog.LogLevel >= logger.Warn {
		glog.logger.Warnf(warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (glog *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if glog.LogLevel >= logger.Error {
		glog.logger.Errorf(errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (glog *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	if glog.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && glog.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !glog.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			glog.logger.Errorf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			glog.logger.Errorf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > glog.SlowThreshold && glog.SlowThreshold != 0 && glog.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", glog.SlowThreshold)
		if rows == -1 {
			glog.logger.Warnf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			glog.logger.Warnf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case glog.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			glog.logger.Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			glog.logger.Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}

// InitLogger 日志类初始化方法
func InitLogger() {
	newZapLogger()
	// gorm logger
	newGormLogger()
}

// 日志记录地址
func getLogWriter() zapcore.WriteSyncer {
	//定义日志文件名，设置权限，当日志文件不存在时创建文件
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.GlobalConfig.Log.Path,
		MaxSize:    1,  // 切割大小
		MaxBackups: 5,  // 保留最大数量
		MaxAge:     30, // 保留最大天数
		Compress:   false,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)

}

// 日志编码方式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// TODO 生产环境 和 开发环境
	//return zapcore.NewJSONEncoder(encoderConfig)
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func newZapLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.Level(config.GlobalConfig.Log.LogLevel))
	zapLogger := zap.New(core, zap.AddCaller())
	Logger = zapLogger.Sugar()
}
