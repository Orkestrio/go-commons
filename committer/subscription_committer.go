package committer

import (
	"github.com/Orkestrio/go-commons/bus"
)

var ACTION_CREATE = "CREATE"
var ACTION_UPDATE = "UPDATE"
var ACTION_DELETE = "DELETE"

type Notification struct {
	Model  string
	Action string
	ID     int
}

type SubscriptionCommitter struct {
	Notifications []Notification
}

func (comm *SubscriptionCommitter) Init() {
	comm.Notifications = []Notification{}
}

func (comm *SubscriptionCommitter) Notify(model string, id int, action string) {
	found := false
	for _, entry := range comm.Notifications {
		if (entry.Model == model) && (entry.ID == id) && (entry.Action == action) {
			found = true
		}
	}

	if !found {
		comm.Notifications = append(comm.Notifications, Notification{Model: model, ID: id, Action: action})
	}
}

func (comm *SubscriptionCommitter) Commit() {
	for _, notification := range comm.Notifications {
		bus.MessageBus.Publish(notification.Model, notification)
	}
}
