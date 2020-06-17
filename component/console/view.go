package console

import (
	"github.com/jroimartin/gocui"
)

// View is a view
func View(g *gocui.Gui, x0, y0, x1, y1 int) error {
	if v, err := g.SetView("console", x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "console"
		v.Autoscroll = true
		v.Wrap = true
	}

	return nil
}
