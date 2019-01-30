package rules

// A Loader loads rules
type Loader interface {
	LoadRules() ([]*Rule, error)
}

// A Saver saves rules
type Saver interface {
	SaveRules([]*Rule) error
}
