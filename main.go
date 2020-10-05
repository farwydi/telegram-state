package main

import (
	"context"
	"github.com/farwydi/cleanwhale/config"
	"github.com/farwydi/cleanwhale/log"
	"github.com/farwydi/cleanwhale/metrics"
	"github.com/farwydi/cleanwhale/transport/whalehttp"
	"github.com/farwydi/cleanwhale/wave"
	"telegram-state/domain"
)

func main() {
	var cfg domain.Config
	err := config.LoadConfigs(&cfg, "config.yml")
	if err != nil {
		log.Fatal(err)
	}

	logger, err := log.NewLogger(cfg.Project)
	if err != nil {
		log.Fatal(err)
	}

	mlog := logger.Named("main")

	metrics.RegisterMetrics("", mlog)

	w := wave.NewWave(context.Background(), mlog)

	w.AddSever(whalehttp.NewHTTPServer(
		cfg.Transport.Webhook, logger.Named("webhook"), nil))

	w.Run()
}
