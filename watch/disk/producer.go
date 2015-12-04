package disk

import (
	"log"
	"time"

	"github.com/erichnascimento/cloud-watch/config"
)

type Producer struct {
	interval *time.Ticker
	disks    []config.Disk
	target   chan<- *Info
}

func NewProducer(interval *time.Ticker, disks []config.Disk) *Producer {
	return &Producer{
		interval: interval,
		disks:    disks,
	}
}

func (p *Producer) Start(target chan<- *Info) {
	p.target = target

	p.Produce()
	go func() {
		for {
			<-p.interval.C
			p.Produce()
		}
	}()
}

func (p *Producer) Produce() {
	for _, d := range p.disks {
		info, err := CollectDiskInfo(d.Path, d.Label)
		if err != nil {
			log.Fatalf("error collecting disk (%s) info: %s", d.Label, err)
		}

		if info.UsedPercentage() >= d.Threshold {
			p.target <- info
		}
	}
}
