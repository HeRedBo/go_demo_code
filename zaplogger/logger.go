package zaplogger

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var Logger *zap.Logger

// InitLogger 初始化全局Logger
func InitLogger(logPath string, logLevel zapcore.Level) error {
	// 1. 设置日志文件路径和轮转规则
	// 例如：在 ./logs 目录下生成以日期命名的日志文件，保留最近7天的日志
	fileWriter, err := rotatelogs.New(
		path.Join(logPath, "app.%Y%m%d.log"),      // 轮转后的文件名模式
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间（7天）
		rotatelogs.WithRotationTime(24*time.Hour), // 日志轮转间隔（24小时）
	)
	if err != nil {
		return err
	}
	// 2. 核心：自定义编码器配置
	fileEncoderConfig := zap.NewProductionEncoderConfig()
	// 关键修改：自定义文件日志的时间格式
	// 使用 Go 的“神奇时间点”来定义格式：2006-01-02 15:04:05.000
	fileEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	// 你也可以取消下面这行的注释，使用 ISO8601 格式（带时区）
	// fileEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 或者，如果你想时间更紧凑：
	// fileEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05")
	fileEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 级别用大写，如 INFO
	// --- 用于控制台的Console编码器配置 ---
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	// 为控制台设置一个更简洁、不带日期的时间格式
	//consoleEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05.000")
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 带颜色的级别
	// 3. 配置 Zap 的编码器和输出端
	// 同时输出到文件和标准输出（控制台）
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel // WARN及以上级别的日志
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logLevel && lvl < zapcore.WarnLevel // 配置级别到INFO的日志
	})
	// 控制台输出编码器（适合人类阅读）
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)
	// 文件输出编码器（JSON格式，适合机器解析）
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)
	// 核心（Core）：将编码器、输出目的地、日志级别绑定
	core := zapcore.NewTee(
		// 将INFO和DEBUG级别的日志输出到文件（JSON格式）
		zapcore.NewCore(fileEncoder, zapcore.AddSync(fileWriter), lowPriority),
		// 将WARN及以上级别的日志同时输出到文件和控制台
		zapcore.NewCore(fileEncoder, zapcore.AddSync(fileWriter), highPriority),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), highPriority),
	)
	// 4. 创建Logger，并添加调用者信息（文件名和行号）
	Logger = zap.New(core, zap.AddCaller())
	return nil
}

func GetLogger() *zap.Logger {
	return Logger
}
