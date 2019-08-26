package main

import (
	"flag"
	"log"
	"time"

	"github.com/aleiphoenix/es-cleaner/pkg/config"
	"github.com/aleiphoenix/es-cleaner/pkg/es"
	"github.com/mileusna/crontab"
)

func main() {
	var dry = flag.Bool("dry", true, "dry mode")
	var configFile = flag.String("config", "config.yaml", "config file")
	flag.Parse()

	log.Printf("dry mode: %v", *dry)

	cfg := config.ReadConfig(*configFile)

	cron := crontab.New()
	cron.MustAddJob("* * * * *", CleanLog, cfg, *dry)

	for {
		time.Sleep(time.Second * 10)
	}

}

func CleanLog(cfg *config.Config, dry bool) {

	log.Println("CleanLog")

	for _, job := range cfg.CleanJob {
		err := es.CleanByDay(
			job.Endpoints,
			job.NamePattern,
			job.DatePattern,
			job.Retention,
			dry,
		)
		if err != nil {
			panic(err)
		}
	}

}
