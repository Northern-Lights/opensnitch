package rules

import (
	"github.com/evilsocket/opensnitch/network"
)

// A Matcher matches connections against rules
type Matcher interface {
	Match(*network.Connection) *Rule
}
