package logger

//用zap日志库实现日志
import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"honnef.co/go/tools/config"
)

//全局logger
var logger *zap.Logger

//Init 初始化logger
func Init(cfg *config.LogConfig) (err error) {
	//来自于config.yaml的配置
	writeSyncer := getWriteSyncer(
		cfg.Filename,
		cfg.Maxsize,
		cfg.MaxBackups,
		cfg.MaxAge)
	encoder := getEncoder(cfg)
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, *l)
	logger = zap.New(core, zap.AddCaller())
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
