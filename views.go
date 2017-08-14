package main

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/willf/pad"
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

// View: Pods
func viewPods(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("pods", -1, 1, lMaxX, lMaxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Settings
		v.Frame = false
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.SetCursor(0, 2)

		// Set as current view
		g.SetCurrentView(v.Name())

		// Content
		// Add column
		viewPodsAddLine(v, lMaxX, "NAME", "READY", "STATUS", "RESTARTS", "AGE")
		fmt.Fprintln(v, strings.Repeat("─", lMaxX))
	}

	return nil
}

// Add line to view pods
func viewPodsAddLine(v *gocui.View, maxX int, name, ready, status, restarts, age string) {
	wN := maxX - 40 - 2
	if wN < 45 {
		wN = 45
	}
	line := pad.Right(name, wN, " ") +
		pad.Right(ready, 12, " ") +
		pad.Right(status, 12, " ") +
		pad.Right(restarts, 12, " ") +
		pad.Right(age, 4, " ")
	fmt.Fprintln(v, line)
}
