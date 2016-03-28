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
	m.diskInfoQueue = make(chan *disk.Info, len(c.Disks))

	m.dispatcher = notification.NewDispatcher(c.Notification, m.diskInfoQueue)

	m.diskProducer = disk.NewProducer(c.Disks, m.diskInfoQueue)
}

func (m *Monitor) Start() {
	// starts metric consumers
	m.dispatcher.Start()

	func() {
		for {
			// execute metric producers
			m.exec()
			<-m.interval.C
		}
	}()
}

func (m *Monitor) exec() {
	go m.diskProducer.Produce()
}

func NewMonitor(c *config.Config) *Monitor {
	m := new(Monitor)
	m.Reconfigure(c)

	return m
}
