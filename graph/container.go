package graph

import (
	"os"

	"github.com/caarlos0/env/v6"
	validator "github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ServiceName string = "xiv-craftsmanship-api"
)

type Container struct {
	Logger     *zap.Logger
	Env        *Env
	Validator  *validator.Validate
	PostgreSQL *sqlx.DB
}

// NOTE: 環境変数の構造体
type Env struct {
	Stage              string `env:"STAGE" envDefault:"development"`
	Port               string `env:"PORT" envDefault:"8080"`
	PostgreSqlHost     string `env:"POSTGRE_SQL_HOST" envDefault:"localhost:5432"`
	PostgreSqlUsername string `env:"POSTGRE_SQL_USERNAME" envDefault:"example"`
	PostgreSqlPassword string `env:"POSTGRE_SQL_PASSWORD" envDefault:"example"`
	PostgreSqlDb       string `env:"POSTGRE_SQL_DB" envDefault:"example"`
}

func NewContainer() (*Container, func()) {
	logger := logger()
	env := environment(logger)
	postgresql, disconnect := postgresql(env, logger)

	return &Container{
			Logger:     logger,
			Env:        env,
			Validator:  validator.New(),
			PostgreSQL: postgresql,
		}, func() { // defer func
			defer logger.Sync()
			defer disconnect()
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
		// os.Exit(1)
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

func postgresql(env *Env, logger *zap.Logger) (*sqlx.DB, func()) {
	// データベースの接続文字列
	uri := "postgresql://" +
		env.PostgreSqlUsername +
		":" +
		env.PostgreSqlPassword +
		"@" +
		env.PostgreSqlHost +
		"/" +
		env.PostgreSqlDb +
		"?sslmode=disable"

	// データベースに接続
	db, err := sqlx.Connect("postgres", uri)
	if err != nil {
		logger.Error(ServiceName, zap.String("message", "Failed to connect to database"), zap.Error(err))
		os.Exit(1)
	}

	return db, func() {
		defer db.Close()
	}
}
