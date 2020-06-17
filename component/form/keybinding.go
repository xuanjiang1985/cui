package form

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
)

// Keybinding action
func Keybinding(g *gocui.Gui) error {
	if err := g.SetKeybinding("form", gocui.MouseLeft, gocui.ModNone, showMsg); err != nil {
		return err
	}
	return nil
}

func showMsg(g *gocui.Gui, v *gocui.View) error {

	// g.Update(func(g *gocui.Gui) error {
	// 	v, err := g.View("console")
	// 	if err != nil {
	// 		return err
	// 	}
	// 	//scrollView(v, 1)
	// 	// v.Clear()
	// 	fmt.Fprintln(v, time.Now().Unix())
	// 	return nil
	// })

	v, err := g.View("console")
	if err != nil {
		return err
	}

	fmt.Fprintln(v, time.Now().Unix())
	return nil

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
