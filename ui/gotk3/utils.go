package gotk3

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func getObject(b *gtk.Builder, name string) interface{} {
	obj, err := b.GetObject(name)
	if err != nil {
		log.Fatalf("Couldn't get %s: %v", name, err)
	}
	return obj
}

func getButton(b *gtk.Builder, name string) *gtk.Button {
	return getObject(b, name).(*gtk.Button)
}

func getCheckButton(b *gtk.Builder, name string) *gtk.CheckButton {
	return getObject(b, name).(*gtk.CheckButton)
}

func getDialog(b *gtk.Builder, name string) *gtk.Dialog {
	return getObject(b, name).(*gtk.Dialog)
}

func getLabel(b *gtk.Builder, name string) *gtk.Label {
	return getObject(b, name).(*gtk.Label)
}

func getRadioButton(b *gtk.Builder, name string) *gtk.RadioButton {
	return getObject(b, name).(*gtk.RadioButton)
}
