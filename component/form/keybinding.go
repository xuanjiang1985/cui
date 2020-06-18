package form

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
)

// Keybinding action
func Keybinding(g *gocui.Gui) error {
	if err := g.SetKeybinding("input", gocui.MouseLeft, gocui.ModNone, showMsg); err != nil {
		return err
	}

	if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, updateInput); err != nil {
		return err
	}

	return nil
}

func showMsg(g *gocui.Gui, v *gocui.View) error {

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	v2, err := g.View("console")
	if err != nil {
		return err
	}

	fmt.Fprintln(v2, time.Now().Unix())

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

//updateInput 当按下ENTER键时调用，把输入的内容复制到输出窗口中
func updateInput(g *gocui.Gui, v *gocui.View) error {
	v2, err := g.View("console")
	if err != nil {
		return err
	}
	fmt.Fprintln(v2, "input")
	// if cv != nil && err == nil {
	// 	var p = cv.ReadEditor()
	// 	if p != nil {
	// 		v.Write([]byte("你:"))
	// 		v.Write(append(p, '\n'))
	// 	}
	// 	v.Autoscroll = true
	// }
	// l := len(cv.Buffer())
	// cv.MoveCursor(0-l, 0, true)
	v.Clear()
	v.SetCursor(0, 0)
	return nil
}
