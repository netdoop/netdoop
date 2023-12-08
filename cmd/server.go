package cmd

import (
	"context"

	"github.com/netdoop/netdoop/doopfx"
	"github.com/netdoop/netdoop/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var serverRunCommand = &cobra.Command{
	Use:   "run",
	Short: "NetDoop Server",
	Run: func(cmd *cobra.Command, args []string) {
		doopfx.CreateAndServeApp(
			fx.Invoke(runServer),
		)
	},
}

type RunParams struct {
	fx.In
	Env    *viper.Viper
	Logger *zap.Logger
	Server *server.Server
}

func runServer(lc fx.Lifecycle, params RunParams) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := doRun(params); err != nil {
					params.Logger.Error("do run", zap.Error(err))
				}
			}()
			return nil
		},
	})
}

func doRun(params RunParams) error {
	return nil
}
