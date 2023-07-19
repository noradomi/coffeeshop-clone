package main

import (
	"github.com/noradomi/coffeeshop-clone/cmd/counter/config"
	"go.uber.org/automaxprocs/maxprocs"
	"golang.org/x/exp/slog"
)

func main() {
	// set GOMAXPROCS
	_, err := maxprocs.Set()
	if err != nil {
		slog.Error("failed set max procs", err)
	}

	//ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("failed get config", err)
	}

	slog.Info("> init app", "name", cfg.Name, "version", cfg.Version)
}

//func prepareApp(ctx context.Context, cancel context.CancelFunc, cfg *config.Config, server *grpc.Server) func() {
//
//}
