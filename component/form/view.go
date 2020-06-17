package form

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// View is a view
func View(g *gocui.Gui, x0, y0, x1, y1 int) error {
	if v, err := g.SetView("form", x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "form"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Button 1 - line 1")
		fmt.Fprintln(v, "Button 1 - line 2")
		fmt.Fprintln(v, "Button 1 - line 3")
		fmt.Fprintln(v, "Button 1 - line 4")
	}

	return nil
}
