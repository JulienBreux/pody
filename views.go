package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// View: Overlay
func viewOverlay(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("overlay", 0, 0, lMaxX, lMaxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Settings
		v.Frame = false
	}

	return nil
}

// View: Title bar
func viewTitle(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("title", -1, -1, lMaxX, 1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Settings
		v.Frame = false
		v.BgColor = gocui.ColorDefault | gocui.AttrReverse
		v.FgColor = gocui.ColorDefault | gocui.AttrReverse

		// Content
		fmt.Fprintln(v, versionTitle(lMaxX))
	}

	return nil
}

// View: Debug
func viewDebug(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("debug", 2, 2, lMaxX-4, lMaxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Settings
		v.Title = " Debug "
		v.Autoscroll = true
	}

	return nil
}
