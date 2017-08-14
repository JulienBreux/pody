package main

import "github.com/jroimartin/gocui"

var DEBUG_DISPLAYED bool = false

// Global action: Quit
func actionGlobalQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// Global action: Toggle debug
func actionGlobalToggleDebug(g *gocui.Gui, v *gocui.View) error {
	vn := "debug"

	if !DEBUG_DISPLAYED {
		debug(g, "Action: Toggle debug display (show)")
		g.SetViewOnTop(vn)
		g.SetCurrentView(vn)
	} else {
		debug(g, "Action: Toggle debug display (hide)")
		g.SetViewOnBottom(vn)
		g.SetCurrentView("pods")
	}

	DEBUG_DISPLAYED = !DEBUG_DISPLAYED

	return nil
}
