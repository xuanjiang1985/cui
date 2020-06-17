package console

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// Keybinding action
func Keybinding(g *gocui.Gui) error {
	if err := g.SetKeybinding("console", gocui.KeyCtrlX, gocui.ModNone, clearMsg); err != nil {
		return err
	}
	return nil
}

func clearMsg(g *gocui.Gui, v *gocui.View) error {
	v, err := g.View("console")
	if err != nil {
		return err
	}

	scrollView(v, 5)

	fmt.Fprintln(v, "clear")

	return nil
}

func scrollView(v *gocui.View, dy int) error {
	if v != nil {
		v.Autoscroll = false
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy+dy); err != nil {
			return err
		}
	}
	return nil
}
