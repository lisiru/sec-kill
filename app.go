package api_server_demo

import (
	"sec-kill/config"
	"sec-kill/options"
	"sec-kill/pkg/app"
	"sec-kill/pkg/logger"
	"sec-kill/pkg/server"
	"os"
)

func NewApp() error  {
	opts:=options.NewOptions()
	if err:=app.AddConfigToOptions(opts);err!=nil {
		os.Exit(1)
	}

	logger.Init(opts.Log)
	defer logger.Flush()
	cfg,err:=config.CreateConfigFromOptions(opts)
	if err!=nil {
		return err
	}
	stopCh:=server.SetupSignalHandler()
	return Run(cfg,stopCh)
}
