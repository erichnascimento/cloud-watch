package notification

import (
	"fmt"
	"log"

	"github.com/erichnascimento/cloud-watch/config"
	"github.com/erichnascimento/cloud-watch/notification/email"
	"github.com/erichnascimento/cloud-watch/watch/disk"
)

type Dispatcher struct {
	config   config.Notification
	diskInfo chan *disk.Info
	mailer   *email.Mailer
}

func NewDispatcher(c config.Notification) *Dispatcher {
	return &Dispatcher{c, make(chan *disk.Info), email.NewMailer()}
}

func (d *Dispatcher) Start() chan<- *disk.Info {
	go func() {
		for {
			select {
			case info := <-d.diskInfo:
				d.sendDiskNotification(info)
			}
		}
	}()

	return d.diskInfo
}

func (d *Dispatcher) sendDiskNotification(info *disk.Info) {
	content := fmt.Sprintf(`
The disk usage limit may be reached:
  Disk:  %s(%s)
  Usage: %.2f%%
`, info.Label, info.Path, info.UsedPercentage())

	log.Printf("DANGER: Disk %s (%s) Usage: %.2f%%", info.Label, info.Path, info.UsedPercentage())
	//fmt.Print(content)
	err := d.mailer.SendMail(content)
	if err != nil {
		log.Fatalf("error sending email: %s", err)
	}
}
