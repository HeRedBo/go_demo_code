package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main01() {
	// 创建生产环境 Logger（JSON 格式）
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// 记录结构化日志
	logger.Info("User login",
		zap.String("username", "alice"),
		zap.Int("attempt", 3),
	)

	// 创建开发环境 Logger（控制台格式）
	devLogger, _ := zap.NewDevelopment()
	devLogger.Debug("Debug message") // 仅开发环境输出
}

func InitLogger() *zap.Logger {
	// 1. 配置文件输出
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/Users/hehongbo/www/GO/go_demo_code/pack/app.log", // 日志文件路径
		MaxSize:    100,                                                // 单文件最大100MB
		MaxBackups: 3,                                                  // 保留3个备份
		MaxAge:     30,                                                 // 保留30天
	})

	// 2. 配置控制台输出
	consoleWriter := zapcore.Lock(os.Stdout) // 并发安全

	// 3. 定义编码器（JSON格式）
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	//encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 4. 创建两个核心（文件和控制台）
	fileCore := zapcore.NewCore(encoder, fileWriter, zap.InfoLevel)
	consoleCore := zapcore.NewCore(encoder, consoleWriter, zap.DebugLevel)

	// 5. 合并核心
	core := zapcore.NewTee(fileCore, consoleCore)

	// 6. 创建 Logger
	return zap.New(core, zap.AddCaller())
}

func main() {
	logger := InitLogger()
	defer logger.Sync()

	logger.Info("this is info msg")
	logger.Warn("this is warn msg")
	//logger.Error("this is error msg")
}

func InitLogger2() *zap.Logger {
	// 配置日志切割
	lumberJack := &lumberjack.Logger{
		Filename:   "/Users/hehongbo/www/GO/go_demo_code/pack/app.log",
		MaxSize:    100, // MB
		MaxBackups: 5,
		MaxAge:     30, // 天
	}

	// 创建核心配置
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(lumberJack),
		zap.InfoLevel,
	)

	return zap.New(core, zap.AddCaller())
}
