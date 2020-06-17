package help

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// View is a view
func View(g *gocui.Gui, x0, y0, x1, y1 int) error {
	if v, err := g.SetView("help", x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "help"
		fmt.Fprintln(v, "ctrl + c: Exit; ctrl + x: clear console")
	}

	return nil
}
