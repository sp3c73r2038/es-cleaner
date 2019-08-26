package main

import (
	"flag"
	"log"

	"github.com/aleiphoenix/es-cleaner/pkg/config"
	"github.com/aleiphoenix/es-cleaner/pkg/es"
)

func main() {
	var dry = flag.Bool("dry", true, "dry mode")
	var configFile = flag.String("config", "config.yaml", "config file")
	flag.Parse()

	log.Printf("dry mode: %v", *dry)

	cfg := config.ReadConfig(*configFile)

	for _, job := range cfg.CleanJob {
		CleanLog(job, *dry)
	}

}

func CleanLog(job config.CleanJob, dry bool) {

	log.Println("CleanLog")
	log.Printf("job: %v", job)
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
