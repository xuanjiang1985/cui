package component

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// Pipe accept msg
func Pipe(g *gocui.Gui, done chan struct{}, msg chan string) {
	for {
		select {
		case <-done:
			return
		case a := <-msg:

			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("console")
				if err != nil {
					return err
				}
				fmt.Fprintln(v, a)
				return nil
			})
		}
	}
}
