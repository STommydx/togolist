package app

import (
	"github.com/STommydx/togolist/internal/pkg/config"
	"github.com/STommydx/togolist/internal/pkg/db"
	"github.com/STommydx/togolist/internal/pkg/grpc"
	"github.com/STommydx/togolist/internal/pkg/healthz"
	"github.com/STommydx/togolist/internal/pkg/http"
	"github.com/STommydx/togolist/internal/pkg/list"
	"github.com/STommydx/togolist/internal/pkg/logger"
	"github.com/STommydx/togolist/internal/pkg/router"
	"go.uber.org/fx"
)

func New() *fx.App {
	app := fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(db.New),
		logger.Module,
		http.Module,
		router.Module,
		healthz.Module,
		list.Module,
		grpc.Module,
	)
	return app
}
