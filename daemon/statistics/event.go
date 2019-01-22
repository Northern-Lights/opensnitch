package statistics

import (
	"time"

	"github.com/Northern-Lights/os-rules-engine/network"
	"github.com/Northern-Lights/os-rules-engine/rules"
	protocol "github.com/Northern-Lights/os-rules-engine/ui"
)

type Event struct {
	Time       time.Time
	Connection *network.Connection
	Rule       *rules.Rule
}

func NewEvent(con *network.Connection, match *rules.Rule) *Event {
	return &Event{
		Time:       time.Now(),
		Connection: con,
		Rule:       match,
	}
}

func (e *Event) Serialize() *protocol.Event {
	return &protocol.Event{
		Time:       e.Time.Format("2006-01-02 15:04:05"),
		Connection: e.Connection,
		Rule:       e.Rule,
	}
}
