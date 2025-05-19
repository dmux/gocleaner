package cmd

import (
	"fmt"
	"time"

	"gocleaner/config"
	"gocleaner/core"

	"github.com/go-co-op/gocron"
)

func Execute() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	if cfg.Schedule.Enabled {
		scheduler := gocron.NewScheduler(time.Local)
		_, err := scheduler.Cron(cfg.Schedule.Cron).Do(func() {
			fmt.Println("🧹 Executando limpeza agendada...")
			core.RunCleaner(cfg)
		})
		if err != nil {
			panic(err)
		}
		scheduler.StartBlocking()
	} else {
		core.RunCleaner(cfg)
	}
}
