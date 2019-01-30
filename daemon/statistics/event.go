package statistics

import (
	"time"

	"github.com/evilsocket/opensnitch/network"
	"github.com/evilsocket/opensnitch/rules"
	protocol "github.com/evilsocket/opensnitch/ui"
)

const fmtTime = "2006-01-02 15:04:05"

func NewEvent(con *network.Connection, match *rules.Rule) *protocol.Event {
	return &protocol.Event{
		Time:       time.Now().Format(fmtTime),
		Connection: con,
		Rule:       match,
	}
}
