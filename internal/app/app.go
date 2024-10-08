package app

import (
	"context"
	"database/sql"

	"github.com/core-go/health/echo"
	s "github.com/core-go/health/sql"
	"github.com/core-go/log/zap"

	"go-service/internal/user"
)

type ApplicationContext struct {
	Health  *echo.Handler
	Handler user.UserTransport
}

func NewApp(ctx context.Context, cfg Config) (*ApplicationContext, error) {
	db, err := sql.Open(cfg.Sql.Driver, cfg.Sql.DataSourceName)
	if err != nil {
		return nil, err
	}
	logError := log.LogError

	userHandler, err := user.NewUserHandler(db, logError)
	if err != nil {
		return nil, err
	}

	sqlChecker := s.NewHealthChecker(db)
	healthHandler := echo.NewHandler(sqlChecker)

	return &ApplicationContext{
		Health:  healthHandler,
		Handler: userHandler,
	}, nil
}
