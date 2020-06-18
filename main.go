package main

import (
	"cui/component"
	"cui/component/console"
	"cui/component/form"
	"cui/component/help"
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

var (
	done = make(chan struct{})
	msg  = make(chan string, 100)
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Mouse = true
	g.Cursor = true
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

	if err := help.Keybinding(g); err != nil {
		fmt.Println(err)
	}

	go component.Pipe(g, done, msg)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// form component
	formY2 := 12
	if err := form.View(g, 2, 2, maxX-2, formY2); err != nil {
		log.Panicln(err)
	}

	// console component
	if err := console.View(g, 2, formY2+2, maxX-2, maxY-4); err != nil {
		log.Panicln(err)
	}

	// help component
	if err := help.View(g, 0, maxY-3, maxX-1, maxY-1); err != nil {
		log.Panicln(err)
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	close(done)
	return gocui.ErrQuit
}
