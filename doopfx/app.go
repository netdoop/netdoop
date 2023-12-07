package doopfx

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/netdoop/netdoop/utils"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func CreateAndServeApp(opts ...fx.Option) {
	opts2 := []fx.Option{
		Module,
	}
	opts2 = append(opts2, opts...)
	app := CreateAndStartApp(opts2...)
	CatchSignal(app)
}

func CreateAndStartApp(opts ...fx.Option) *fx.App {
	ctx := context.Background()
	logger := utils.GetLogger()
	opts = append(opts, fx.WithLogger(func() fxevent.Logger {
		if utils.DebugMode {
			return &fxevent.ZapLogger{Logger: logger}
		} else {
			return fxevent.NopLogger
		}
	}))
	app := fx.New(opts...)
	if err := app.Start(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return app
}

func StopApp(app *fx.App) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Stop(ctx); err != nil {
		fmt.Println("error from shutdown", err)
	}
}

func CatchSignal(app *fx.App) {
	//terminateSignals := make(chan os.Signal, 1)
	reloadSignals := make(chan os.Signal, 1)

	//signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	signal.Notify(reloadSignals, syscall.SIGUSR1)
	for {
		select {
		case <-app.Done():
			fmt.Println("done")
			StopApp(app)
			return
		//case s := <-terminateSignals:
		//fmt.Println("got stop signals, shutting down server gracefully, SIGNAL:", s)
		//ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		//defer cancel()
		//if err := app.Stop(ctx); err != nil {
		//fmt.Println("error from shutdown", err)
		//return
		//} else {
		//fmt.Println("shutdown")
		////os.Exit(1)
		//}
		case <-reloadSignals:
			fmt.Println("got reload signal, will reload")
			//Some config reload code here.
		}
	}
}
