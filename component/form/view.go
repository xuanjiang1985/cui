package form

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// type signup struct {
// 	*component.Form
// }

// View is a view
func View(g *gocui.Gui, x0, y0, x1, y1 int) error {
	// new form
	// signup := &signup{
	// 	component.NewForm(g, "form", x0, y0, x1, y1),
	// }

	// // add input field
	// signup.AddInputField("First Name", 11, 28).
	// 	AddValidate("required input", requireValidator)
	// signup.AddInputField("Last Name", 11, 28).
	// 	AddValidate("required input", requireValidator)

	// signup.AddInputField("Password", 11, 28).
	// 	AddValidate("required input", requireValidator).
	// 	SetMask().
	// 	SetMaskKeybinding(gocui.KeyCtrlA)

	// // // add checkbox
	// // signup.AddCheckBox("Age 18+", 11)

	// // // add select
	// // signup.AddSelect("Language", 11, 10).AddOptions("Japanese", "English", "Chinese")

	// // // add radios
	// // signup.AddRadio("Country", 11).
	// // 	SetMode(component.VerticalMode).
	// // 	AddOptions("Japan", "America", "China")

	// // add button
	// signup.AddButton("Regist", showMsg)
	// signup.AddButton("Cancel", cancelMsg)

	// signup.Draw()

	// if v, err := g.SetView("form", x0, y0, x1, y1); err != nil {
	// 	if err != gocui.ErrUnknownView {
	// 		return err
	// 	}
	// 	v.Title = "form"
	// 	v.Highlight = true
	// 	v.SelBgColor = gocui.ColorGreen
	// 	v.SelFgColor = gocui.ColorBlack
	// }

	if v, err := g.SetView("input", x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "input"
		v.Editable = true
		v.Wrap = true
		v.Overwrite = false

	}

	return nil
}

func cancelMsg(g *gocui.Gui, v *gocui.View) error {

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	v2, err := g.View("console")
	if err != nil {
		return err
	}

	fmt.Fprintln(v2, "cancel")

	return nil
}

func requireValidator(value string) bool {
	if value == "" {
		return false
	}
	return true
}
