package rule

import (
	"fmt"
	"sync"

	"github.com/Northern-Lights/os-rules-engine/network"
	"github.com/Northern-Lights/os-rules-engine/rules"
)

// Manager manages rules in memory and their storage
type Manager struct {
	sync.RWMutex

	Loader
	Saver

	rules []*rules.Rule
}

func NewManager(opts ...ManagerOption) (*Manager, error) {
	var m Manager
	for _, opt := range opts {
		err := opt(&m)
		if err != nil {
			return nil, err
		}
	}
	return &m, nil
}

// Add adds the given rule and stores it, if a saver is set
func (m *Manager) Add(r *rules.Rule) (err error) {
	m.Lock()
	defer m.Unlock()

	m.rules = append(m.rules, r)

	return
}

// AddAndSave adds the given rule and stores it, if a saver is set
func (m *Manager) AddAndSave(r *rules.Rule) (err error) {
	m.Add(r)
	if m.Saver != nil {
		err = m.Saver.SaveRules(m.rules)
	}

	return
}

// LoadRules uses the Loader to load and replace the in-memory rules
func (m *Manager) LoadRules() (ruleset []*rules.Rule, err error) {
	m.Lock()
	defer m.Unlock()

	if m.Loader == nil {
		err = fmt.Errorf(`rule: no loader set; cannot load rules`)
		return
	}

	ruleset, err = m.Loader.LoadRules()
	if err != nil {
		ruleset = nil
		return
	}

	return
}

func (m *Manager) Count() int {
	m.RLock()
	defer m.RUnlock()

	return len(m.rules)
}

// Match finds and returns the first matching rule or nil if no matches occured
func (m *Manager) Match(conn *network.Connection) *rules.Rule {
	m.RLock()
	defer m.RUnlock()

	for _, r := range m.rules {
		expr, err := DeserializeExpression(r.Condition)
		if err != nil {
			continue
		}

		matched := expr.Evaluate(conn)
		if matched {
			return r
		}
	}

	return nil
}

// A ManagerOption can be used to create a new manager with customizations
type ManagerOption func(m *Manager) error

func WithLoader(loader Loader) ManagerOption {
	return func(m *Manager) error {
		m.Loader = loader
		return nil
	}
}

func WithSaver(saver Saver) ManagerOption {
	return func(m *Manager) error {
		m.Saver = saver
		return nil
	}
}
