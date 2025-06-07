package app

import (
	"github.com/fsnotify/fsnotify"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
)

func (app *App) initConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		app.err = multierror.Append(app.err, err)
	}

	app.config = viper.GetViper()
	app.config.WatchConfig()
	app.config.OnConfigChange(func(in fsnotify.Event) {
		app.logger.Info("config changed")
	})
}
