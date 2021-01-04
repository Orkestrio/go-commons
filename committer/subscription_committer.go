package committer

import (
	"github.com/Orkestrio/go-commons/bus"
)

type Notification struct {
	Model string
	ID    int
}

type SubscriptionCommitter struct {
	Notifications []Notification
}

func (comm *SubscriptionCommitter) Init() {
	comm.Notifications = []Notification{}
}

func (comm *SubscriptionCommitter) Notify(model string, id int) {
	found := false
	for _, entry := range comm.Notifications {
		if (entry.Model == model) && (entry.ID == id) {
			found = true
		}
	}

	if !found {
		comm.Notifications = append(comm.Notifications, Notification{Model: model, ID: id})
	}
}

func (comm *SubscriptionCommitter) Commit() {
	for _, notification := range comm.Notifications {
		bus.MessageBus.Publish(notification.Model, notification.ID)
	}
}
