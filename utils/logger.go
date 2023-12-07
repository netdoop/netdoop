package utils

import (
	"fmt"
	"os"
	"sync"

	"github.com/labstack/gommon/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var defaultLogger *zap.Logger
var defaultLoggerOnce sync.Once
var defaultColorer *color.Color

func GetLogger() *zap.Logger {
	defaultLoggerOnce.Do(initLogger)
	return defaultLogger
}
func GetColorer() *color.Color {
	return defaultColorer
}

func initLogger() {
	v := GetEnv()
	v.SetDefault("logger_level", "info")
	v.SetDefault("logger_format", "json")
	v.SetDefault("logger_file_path", "")
	//v.SetDefault("logger_file_path", "logs")
	v.SetDefault("logger_file_compress", 1)
	if DebugMode {
		v.Set("logger_format", "develop")
		v.Set("logger_file_path", "")
		v.Set("logger_file_compress", 0)
	}
	v.SetDefault("logger_file_maxsize", 100)
	v.SetDefault("logger_file_maxage", 30)
	v.SetDefault("logger_file_maxbackups", 5)

	var (
		level         zapcore.Level
		encoderConfig zapcore.EncoderConfig
		encoder       zapcore.Encoder
		writeSyncer   zapcore.WriteSyncer
	)

	levelText := v.GetString("logger_level")
	if VerboseMode {
		levelText = "debug"
	}
	if err := level.UnmarshalText([]byte(levelText)); err != nil {
		level = zapcore.InfoLevel
	}

	formatText := v.GetString("logger_format")
	switch formatText {
	case "develop":
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	case "json":
		fallthrough
	default:
		encoderConfig = zap.NewProductionEncoderConfig()
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.Timekey = "time"
	encoder = zapcore.NewConsoleEncoder(encoderConfig)

	filePathText := v.GetString("logger_file_path")
	if filePathText != "" {
		lumberjackLogger := &lumberjack.Logger{
			Filename:   filePathText,
			MaxSize:    v.GetInt("logger_file_maxsize"),
			MaxAge:     v.GetInt("logger_file_maxage"),
			MaxBackups: v.GetInt("logger_file_maxbackups"),
			Compress:   v.GetString("logger_file_compress") == "true",
		}
		writeSyncer = zapcore.AddSync(lumberjackLogger)
	} else {
		writeSyncer = zapcore.AddSync(os.Stderr)
	}

	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core)
	if DebugMode {
		logger = logger.WithOptions(zap.AddCaller())
	}
	defaultLogger = logger
	defaultColorer = color.New()

	zap.ReplaceGlobals(logger)

	if DebugMode {
		fmt.Println("debug mode")
	}
	if VerboseMode {
		fmt.Println("verbose mode")
	}
}
