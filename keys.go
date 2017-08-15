package main

import "github.com/jroimartin/gocui"

type Key struct {
	viewname string
	key      interface{}
	handler  func(*gocui.Gui, *gocui.View) error
}

// Define UI key bindings
func uiKey(g *gocui.Gui) error {
	for _, key := range keys {
		if err := g.SetKeybinding(key.viewname, key.key, gocui.ModNone, key.handler); err != nil {
			return err
		}
	}

	return nil
}
