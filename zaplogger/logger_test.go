package zaplogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestLogger(t *testing.T) {

	err := InitLogger("./logs", zapcore.InfoLevel)
	if err != nil {
		panic(err) // 或者用其他方式处理初始化失败
	}
	defer GetLogger().Sync() // 程序退出前刷新缓冲区的日志
	log := GetLogger()

	// 记录不同级别的日志
	log.Info("应用程序启动",
		zap.String("version", "1.0.0"),
		zap.Int("port", 8080),
	)

	log.Debug("这是一条Debug日志，因为级别是Info，所以这条不会输出到文件") // 不会出现

	log.Warn("磁盘空间不足",
		zap.String("path", "/data"),
		zap.Float64("free_percent", 15.5),
	)

	log.Error("处理请求失败",
		zap.String("url", "/api/user"),
		zap.Int("status_code", 500),
	)

	// 使用Sugar()方法可以获得更简洁的语法（性能稍差）
	sugarLog := log.Sugar()
	sugarLog.Infof("用户 %s 登录成功，来自IP：%s", "Alice", "192.168.1.100")

}
