package log

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Sugar *zap.SugaredLogger

func Init() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	Sugar = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file := fmt.Sprintf("%s/log_%s.log", "/logs", carbon.Now(carbon.Shanghai).Now().ToDateString())
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    viper.GetInt("LOG_MAX_SIZE"),
		MaxBackups: viper.GetInt("MAX_BACKUPS"),
		MaxAge:     viper.GetInt("LOG_MAX_AGE"),
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
