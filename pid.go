package main

import (
	"fmt"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/relistan/go-director"
	"github.com/shirou/gopsutil/process"
)

const (
	WATCH_INTERVAL = time.Duration(1) * time.Second
)

type PID struct {
	Process *process.Process
	Looper  director.Looper
	stats   []*StatsEntry
	mutex   *sync.Mutex
}

type StatsEntry struct {
	Time    time.Time
	VMS     uint64
	RSS     uint64
	Swap    uint64
	CPU     float64
	Threads int32
}

func NewPID(pid int32) (*PID, error) {
	proc, err := process.NewProcess(pid)
	if err != nil {
		return nil, err
	}

	_, err = proc.Status()
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch initial process status (is the PID running?)")
	}

	return &PID{
		Process: proc,
		Looper:  director.NewTimedLooper(director.FOREVER, WATCH_INTERVAL, nil),
		mutex:   &sync.Mutex{},
	}, nil
}

func (p *PID) Watch() {
	p.Looper.Loop(func() error {
		log.Debugf("Fetching state for pid %v", p.Process.Pid)

		_, err := p.Process.Status()
		if err != nil {
			log.Errorf("Unable to fetch process status (did the proces exit?): %v", err)
			return nil
		}

		statsEntry, err := p.generateStatsEntry()
		if err != nil {
			log.Warningf("Unable to collect pid stats: %v", err)
			return nil
		}

		p.updateStats(statsEntry)
		return nil
	})
}

func (p *PID) updateStats(entry *StatsEntry) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.stats = append(p.stats, entry)
}

func (p *PID) GetStats() []StatsEntry {
	copyStats := make([]StatsEntry, len(p.stats))
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, v := range p.stats {
		copyStats = append(copyStats, *v)
	}

	return copyStats
}

func (p *PID) generateStatsEntry() (*StatsEntry, error) {
	// process is running
	meminfo, err := p.Process.MemoryInfo()
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch memory info: %v", err)
	}

	percent, err := p.Process.Percent(0)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch CPU usage info: %v", err)
	}

	threads, err := p.Process.NumThreads()
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch thread count: %v", err)
	}

	return &StatsEntry{
		RSS:     meminfo.RSS,
		VMS:     meminfo.VMS,
		Swap:    meminfo.Swap,
		CPU:     percent,
		Threads: threads,
		Time:    time.Now(),
	}, nil
}
