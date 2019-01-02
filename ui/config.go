package ui

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/evilsocket/opensnitch/daemon/rule"
)

func init() {
	Config = config{
		File:     os.ExpandEnv(DefaultConfigPath),
		Timeout:  DefaultConfigTimeout,
		Action:   DefaultConfigAction,
		Duration: DefaultConfigDuration,
	}
}

// Config defaults
const (
	DefaultConfigPath     = "$HOME/.opensnitch/ui-config.json"
	DefaultConfigTimeout  = 15 * time.Second
	DefaultConfigAction   = rule.Allow
	DefaultConfigDuration = rule.Restart
)

var configLock sync.RWMutex

// Config is the global configuration
var Config config

type config struct {
	File     string        `json:"-"`
	Timeout  time.Duration `json:"default_timeout"`
	Action   rule.Action   `json:"default_action"`
	Duration rule.Duration `json:"default_duration"`
}

// Load loads the configuration from the given file path
func (c *config) Load() error {
	configLock.Lock()
	defer configLock.Unlock()

	f, err := os.Open(c.File)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(c)
	if err != nil {
		return err
	}
	return nil
}

// Save saves the configuration
func (c *config) Save() error {
	configLock.RLock()
	defer configLock.RUnlock()

	f, err := os.Open(c.File)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	err = enc.Encode(c)
	if err != nil {
		return err
	}
	return nil
}
