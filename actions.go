package main

import "github.com/jroimartin/gocui"

var DEBUG_DISPLAYED bool = false

// Global action: Quit
func actionGlobalQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// Global action: Toggle debug
func actionGlobalToggleDebug(g *gocui.Gui, v *gocui.View) error {
	if !DEBUG_DISPLAYED {
		debug(g, "Action: Toggle debug display (show)")
		g.SetViewOnTop("debug")
	} else {
		debug(g, "Action: Toggle debug display (hide)")
		g.SetViewOnBottom("debug")
		// TODO g.SetCurrentView("pods")
	}

	DEBUG_DISPLAYED = !DEBUG_DISPLAYED

	return nil
}
