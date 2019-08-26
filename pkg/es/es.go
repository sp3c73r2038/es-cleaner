package es

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/olivere/elastic"
)

func CleanByDay(
	endpoints []string,
	namePattern string,
	datePattern string,
	retention int,
	dry bool) error {
	var err error

	tz := "+08:00"
	dp := fmt.Sprintf("%s_-07:00", datePattern)

	cst, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}
	now := time.Now()
	probe := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, cst)
	delta := time.Duration(time.Hour * time.Duration(retention*24))

	client, err := elastic.NewClient(
		elastic.SetURL(endpoints...),
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	indexNames, err := client.IndexNames()
	if err != nil {
		return err
	}

	tmp := make([]string, 0)
	for _, n := range indexNames {
		var m bool
		if m, err = filepath.Match(namePattern, n); err != nil {
			return err
		}
		if m {
			tmp = append(tmp, n)
		}
	}

	targets := make([]string, 0)
	for _, n := range tmp {
		var t time.Time
		nn := fmt.Sprintf("%s_%s", n, tz)
		if t, err = time.Parse(dp, nn); err != nil {
			return err
		}

		if probe.Sub(t) > delta {
			// log.Printf("will delete %s", n)
			targets = append(targets, n)
		}
	}

	for _, n := range targets {
		log.Printf("  delete index %s", n)
	}

	if !dry {
		resp, err := client.DeleteIndex(targets...).Do(context.Background())
		if err != nil {
			return err
		}
		log.Printf("delete request acknowledged: %v", resp.Acknowledged)
	} else {
		log.Println("dry mode, will not actually delete")
	}

	return nil
}
