package rule

import (
	"github.com/Northern-Lights/os-rules-engine/network"
	"github.com/Northern-Lights/os-rules-engine/rules"
)

// A Matcher matches connections against rules
type Matcher interface {
	Match(*network.Connection) *rules.Rule
}
