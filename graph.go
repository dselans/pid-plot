package main

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
)

const (
	MIN_DURATION_SECONDS = float64(10)
)

type XYDataSet struct {
	XValues []time.Time
	YValues []float64
}

// Generate CPU timeseries data set; backfill dataset for time periods with
// no stats data.
func GenerateCPUDataset(duration time.Duration, stats []StatsEntry) (*XYDataSet, error) {
	if duration.Seconds() < MIN_DURATION_SECONDS {
		return nil, fmt.Errorf("Given duration '%v' is less than minimum duration of %v", duration, MIN_DURATION_SECONDS)
	}

	var dataset XYDataSet

	// If our existing stats do not have enough data, fill it with time period and 0's
	backfill := 0
	if len(stats) < int(duration.Seconds()) {
		backfill = int(duration.Seconds()) - len(stats)
		log.Debugf("Backfilling by %v entries", backfill)
	}

	if backfill > 0 {
		for i := backfill; i != 0; i-- {
			dataset.XValues = append(dataset.XValues, stats[0].Time.Add(time.Duration(-i)*time.Second))
			dataset.YValues = append(dataset.YValues, 0)
		}
	}

	// Now let's fill it with real data (with correct startpoint in case we had to backfill)
	start := len(stats) - (int(duration.Seconds()) - backfill)
	for i := start; i != len(stats); i++ {
		dataset.XValues = append(dataset.XValues, stats[i].Time)
		dataset.YValues = append(dataset.YValues, stats[i].CPU)
	}

	return &dataset, nil
}

func GenerateMemDataset(duration time.Duration, stats []StatsEntry) (*XYDataSet, error) {
	if duration.Seconds() < MIN_DURATION_SECONDS {
		return nil, fmt.Errorf("Given duration '%v' is less than minimum duration of %v", duration, MIN_DURATION_SECONDS)
	}

	var dataset XYDataSet

	// If our existing stats do not have enough data, fill it with time period and 0's
	backfill := 0
	if len(stats) < int(duration.Seconds()) {
		backfill = int(duration.Seconds()) - len(stats)
		log.Debugf("Backfilling by %v entries", backfill)
	}

	if backfill > 0 {
		for i := backfill; i != 0; i-- {
			// append oldest time first
			dataset.XValues = append(dataset.XValues, stats[0].Time.Add(time.Duration(-i)*time.Second))
			dataset.YValues = append(dataset.YValues, 0)
		}
	}

	// Now let's fill it with real data (with correct startpoint in case we had to backfill)
	start := len(stats) - (int(duration.Seconds()) - backfill)
	for i := start; i != len(stats); i++ {
		memMB := (float64(stats[i].RSS) / 1024) / 1024 // b -> MB
		dataset.XValues = append(dataset.XValues, stats[i].Time)
		dataset.YValues = append(dataset.YValues, memMB)
	}

	return &dataset, nil
}
