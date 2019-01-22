package rule

import (
	"github.com/Northern-Lights/os-rules-engine/rules"
)

// A Loader loads rules from a reader source
type Loader interface {
	LoadRules() ([]*rules.Rule, error)
}

// A Saver saves rules to a storage writer
type Saver interface {
	SaveRules([]*rules.Rule) error
}
