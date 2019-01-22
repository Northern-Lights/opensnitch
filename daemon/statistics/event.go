package statistics

import (
	"time"

	"github.com/Northern-Lights/os-rules-engine/network"
	"github.com/Northern-Lights/os-rules-engine/rules"
	protocol "github.com/Northern-Lights/os-rules-engine/ui"
)

const fmtTime = "2006-01-02 15:04:05"

func NewEvent(con *network.Connection, match *rules.Rule) *protocol.Event {
	return &protocol.Event{
		Time:       time.Now().Format(fmtTime),
		Connection: con,
		Rule:       match,
	}
}
