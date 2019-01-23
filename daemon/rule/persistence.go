package rule

import (
	"github.com/Northern-Lights/os-rules-engine/rules"
)

// A Loader loads rules
type Loader interface {
	LoadRules() ([]*rules.Rule, error)
}

// A Saver saves rules
type Saver interface {
	SaveRules([]*rules.Rule) error
}
