package gotk3

import (
	"fmt"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var (
	app *gtk.Application
)

// prompt could have been an interface and returned upon init

// Init initializes the application and its components
func Init(appID string) error {
	var err error
	app, err = gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		return err
	}

	var activateErr error
	activate := func() {
		// FIXME: how do we catch this?
		activateErr = initPrompt()

		if activateErr != nil {
			return
		}

		app.AddWindow(&dialog.Window)
	}

	// I think this is synchronous, so activateErr will work...
	_, err = app.Connect("activate", activate)
	if err != nil {
		return err
	}

	if activateErr != nil {
		return activateErr
	}

	return nil
}

// Quit exits the running application
func Quit() error {
	if app == nil {
		return fmt.Errorf("gotk3: application not initialized")
	}

	app.Quit()

	return nil
}

// Run runs the application
func Run() error {
	if app == nil {
		return fmt.Errorf("gotk3: application not initialized")
	}

	go func() {
		app.Run(os.Args)
	}()

	return nil
}
