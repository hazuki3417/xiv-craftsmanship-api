package graph

import (
	"os"

	"github.com/caarlos0/env/v6"
	validator "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ServiceName string = "xiv-craftsmanship-api"
)

type Container struct {
	Logger    *zap.Logger
	Env       *Env
	Validator *validator.Validate
}

// NOTE: 環境変数の構造体
type Env struct {
	Environment string `env:"ENV" envDefault:"develop"`
	Port        string `env:"PORT" envDefault:"8080"`
}

func NewContainer() (*Container, func()) {
	logger := logger()
	env := environment(logger)

	return &Container{
			Logger:    logger,
			Env:       env,
			Validator: validator.New(),
		}, func() { // defer func
			defer logger.Sync()
		}
}

func environment(logger *zap.Logger) *Env {
	// NOTE: .envはローカル開発用のみ利用する想定。.envを環境変数に読み込む。
	// NOTE: .envと実行環境の両方に同じ変数名がある場合、実行環境の環境変数が優先される。
	if err := godotenv.Load(); err != nil {
		logger.Info(ServiceName, zap.String("message", "No .env file found, using environment variables"))
	} else {
		logger.Info(ServiceName, zap.String("message", ".env file found. Loading .env file."))
	}

	envconfig := Env{}
	// NOTE: 環境変数のバリデーション + パース（環境変数を構造体にマッピング）
	if err := env.Parse(&envconfig); err != nil {
		logger.Error(ServiceName, zap.String("message", "Failed to parse environment variables"), zap.Error(err))
		os.Exit(1)
	}
	return &envconfig
}

func logger() *zap.Logger {

	stdoutEncodingConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "service",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	stderrEncodingConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "service",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	stdoutEncoder := zapcore.NewJSONEncoder(stdoutEncodingConfig)
	stderrEncoder := zapcore.NewJSONEncoder(stderrEncodingConfig)

	stdoutCore := zapcore.NewCore(stdoutEncoder, zapcore.Lock(os.Stdout), zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level <= zapcore.InfoLevel
	}))
	stderrCore := zapcore.NewCore(stderrEncoder, zapcore.Lock(os.Stderr), zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.WarnLevel
	}))

	return zap.New(zapcore.NewTee(stdoutCore, stderrCore))
}
