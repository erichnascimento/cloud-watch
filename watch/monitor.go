package watch

import (
	"time"

	"github.com/erichnascimento/cloud-watch/config"
	"github.com/erichnascimento/cloud-watch/notification"
	"github.com/erichnascimento/cloud-watch/watch/disk"
)

type Monitor struct {
	interval      *time.Ticker
	diskInfoQueue chan *disk.Info
	diskProducer  *disk.Producer
	dispatcher    *notification.Dispatcher
}

func (m *Monitor) Reconfigure(c *config.Config) {
	m.interval = time.NewTicker(time.Second * c.Interval)
	m.dispatcher = notification.NewDispatcher(c.Notification)
	m.diskProducer = disk
	m.diskInfoQueue = make(chan *disk.Info, m.diskProducer.TotalDisks())
}

func (m *Monitor) Start() {
	func() {
		for {
			<-m.interval.C
			m.exec()
		}
	}()
}

func (m *Monitor) exec() {
	m.diskProducer.Produce()
}

func NewMonitor(c *config.Config) *Monitor {
	m := new(Monitor)
	m.Reconfigure(c)

	//diskInfo := notification.NewDispatcher(c.Notification).Start()
	//disk.NewProducer(monitor.interval, c.Disks).Start(diskInfo)
	//
	return m
}
