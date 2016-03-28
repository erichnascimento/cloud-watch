package disk

import (
	"log"

	"github.com/erichnascimento/cloud-watch/config"
)

type Producer struct {
	disks  []config.Disk
	target chan<- *Info
}

func NewProducer(disks []config.Disk, target chan<- *Info) *Producer {
	return &Producer{
		target: target,
		disks:  disks,
	}
}

func (p *Producer) Produce() {
	for _, d := range p.disks {
		info, err := CollectDiskInfo(d.Path, d.Label)
		if err != nil {
			log.Printf("error collecting disk (%s) info: %s", d.Label, err)
			return
		}

		if info.UsedPercentage() >= d.Threshold {
			p.target <- info
		}
	}
}
