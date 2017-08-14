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

		// Content: Add column
		//viewPodsAddLine(v, lMaxX, "NAME", "CPU", "MEMORY", "READY", "STATUS", "RESTARTS", "AGE") // TODO CPU + Memory
		viewPodsAddLine(v, lMaxX, "NAME", "READY", "STATUS", "RESTARTS", "AGE")
		fmt.Fprintln(v, strings.Repeat("─", lMaxX))

		// Content: Add lines
		// TODO Use goroutine
		pods, err := getPods()
		if err != nil {
			panic(err.Error())
		}
		if len(pods.Items) > 0 {
			debug(g, fmt.Sprintf("There are %d pods in the cluster", len(pods.Items)))
			for _, pod := range pods.Items {
				n := pod.GetName()
				//c := "?" // TODO CPU + Memory
				//m := "?" // TODO CPU + Memory
				s := columnHelperStatus(pod.Status)
				r := columnHelperRestarts(pod.Status.ContainerStatuses)
				a := columnHelperAge(pod.CreationTimestamp)
				cr := columnHelperReady(pod.Status.ContainerStatuses)
				viewPodsAddLine(v, lMaxX, n, cr, s, r, a)
				//viewPodsAddLine(v, lMaxX, n, c, m, cr, s, r, a) // TODO CPU + Memory
			}
		} else {
			debug(g, "Pods not found.")
		}
	}

	return nil
}

// Add line to view pods
//func viewPodsAddLine(v *gocui.View, maxX int, name, cpu, memory, ready, status, restarts, age string) { // TODO CPU + Memory
func viewPodsAddLine(v *gocui.View, maxX int, name, ready, status, restarts, age string) {
	wN := maxX - 34 // 54 // TODO CPU + Memory
	if wN < 45 {
		wN = 45
	}
	line := pad.Right(name, wN, " ") +
		//pad.Right(cpu, 10, " ") + // TODO CPU + Memory
		//pad.Right(memory, 10, " ") + // TODO CPU + Memory
		pad.Right(ready, 10, " ") +
		pad.Right(status, 10, " ") +
		pad.Right(restarts, 10, " ") +
		pad.Right(age, 4, " ")
	fmt.Fprintln(v, line)
}
