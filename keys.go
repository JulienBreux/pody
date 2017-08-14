package main

import "github.com/jroimartin/gocui"

type Key struct {
	viewname string
	key      interface{}
	handler  func(*gocui.Gui, *gocui.View) error
}

// Configure globale keys
var globalKeys []Key = []Key{
	Key{"", gocui.KeyCtrlC, actionGlobalQuit},
	Key{"", gocui.KeyCtrlD, actionGlobalToggleDebug},
}

// Define UI key bindings
func uiKey(g *gocui.Gui) error {
	// Glboal keys
	for _, key := range globalKeys {
		if err := g.SetKeybinding(key.viewname, key.key, gocui.ModNone, key.handler); err != nil {
			return err
		}
	}

	return nil
}
