package app

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	loan "github.com/davyc98/billing-engine/internal/billing"
	"github.com/davyc98/billing-engine/internal/pkg/pkgsql"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type App struct {
	database     *sql.DB
	queryBuilder pkgsql.GoquBuilder
	validator    *validator.Validate
	logger       *zap.Logger
	router       *httprouter.Router
	httpServer   *http.Server
	closersFn    []func(context.Context) error
	config       *viper.Viper
	err          error
}

func Run() {
	app := &App{}

	app.run()
}

func (app *App) Start() error {
	app.logger.Sugar().Info("starting application")
	go func() {
		app.logger.Sugar().Info("http server listen on", app.httpServer.Addr)
		if err := app.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			app.err = multierror.Append(app.err, err)
		}
	}()

	return nil
}

func (app *App) Stop(ctx context.Context) error {
	for _, closer := range app.closersFn {
		if err := closer(ctx); err != nil {
			app.err = multierror.Append(app.err, err)
		}
	}

	return nil
}

func (app *App) spinUp() *App {
	app.initConfig()
	app.initLogger()
	app.initDB()
	app.setUpGoqu()
	app.initRouter()
	app.makeHTTPServer()
	app.initValidator()
	app.setUpClosers()

	// spin up module
	app.spinUpLoan()

	return app
}

func (app *App) spinUpLoan() {
	loan.New(loan.Dependencies{
		DB:           app.database,
		Logger:       app.logger.Sugar(),
		QueryBuilder: app.queryBuilder,
		HttpRouter:   app.router,
		Validator:    app.validator,
	})
}
