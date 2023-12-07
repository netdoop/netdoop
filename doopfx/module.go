package doopfx

import (
	"context"

	"github.com/netdoop/netdoop/server"
	"github.com/netdoop/netdoop/utils"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"doop",
	fx.Provide(
		NewServer,
	),
)

type ServerResult struct {
	fx.Out
	Env    *viper.Viper
	Logger *zap.Logger
	Server *server.Server
}

func NewServer(lc fx.Lifecycle) (ServerResult, error) {
	env := utils.GetEnv()
	logger := utils.GetLogger()
	s := server.NewServer()
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go s.Run()
			return nil
		},
		OnStop: func(context.Context) error {
			s.Close()
			return nil
		},
	})
	return ServerResult{
		Env:    env,
		Logger: logger,
		Server: s,
	}, nil
}
