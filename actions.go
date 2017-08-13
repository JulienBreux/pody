package main

import "github.com/jroimartin/gocui"

// Global action: Quit
func actionGlobalQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
