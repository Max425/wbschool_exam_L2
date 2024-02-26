package initial

import (
	"context"
	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/pkg/repository"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func InitConfig() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatalf("Ошибка при загрузке файла .env: %s", err)
	}

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func InitLogger() (*zap.Logger, error) {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	return config.Build()
}

func InitPostgres(ctx context.Context) (*sqlx.DB, error) {
	return repository.NewPostgresDB(repository.PostgresConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("PG_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
}
