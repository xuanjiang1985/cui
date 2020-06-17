package help

import (
	"github.com/jroimartin/gocui"
)

// Keybinding action
func Keybinding(g *gocui.Gui) error {
	if err := g.SetKeybinding("help", gocui.MouseLeft, gocui.ModNone, focus); err != nil {
		return err
	}
	return nil
}

func focus(g *gocui.Gui, v *gocui.View) error {

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	return nil
}
