package watch

import (
	"log"
	"time"

	"github.com/erichnascimento/cloud-watch/config"
)

type Monitor struct {
	interval      <-chan time.Time
	config        *config.Config
	diskInfoQueue chan *DiskInfo
}

func (m *Monitor) startProducer() {
	for range m.interval {
		for _, d := range m.config.Disks {
			info, err := CollectDiskInfo(d.Path, d.Label)
			if err != nil {
				log.Fatalf("error collecting disk (%s) info: %s", d.Label, err)
			}
			if info.UsedPercentage() >= d.Threshold {
				m.diskInfoQueue <- info
			}
		}
	}
}

func (m *Monitor) startDiskConsumer() {
	for info := range m.diskInfoQueue {
		log.Printf("DANGER: Disk %s (%s) Usage: %.2f%%", info.Label, info.Path, info.UsedPercentage())
	}
}

func StartMonitor(c *config.Config) error {
	monitor := new(Monitor)
	monitor.config = c
	monitor.interval = time.Tick(time.Second * monitor.config.Interval)
	monitor.diskInfoQueue = make(chan *DiskInfo)

	go monitor.startProducer()
	go monitor.startDiskConsumer()

	return nil
}
