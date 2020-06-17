package main

import (
	"cui/component/console"
	"cui/component/form"
	"cui/component/help"
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := form.Keybinding(g); err != nil {
		log.Panicln(err)
	}

	if err := console.Keybinding(g); err != nil {
		fmt.Println(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// form component
	if err := form.View(g, 2, 2, maxX-2, 7); err != nil {
		log.Panicln(err)
	}

	// console component
	if err := console.View(g, 2, 8, maxX-2, 16); err != nil {
		log.Panicln(err)
	}

	// help component
	if err := help.View(g, 0, maxY-3, maxX-1, maxY-1); err != nil {
		log.Panicln(err)
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
