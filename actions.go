package main

import (
	"github.com/jroimartin/gocui"
)

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

// View pods: Up
func actionViewPodsUp(g *gocui.Gui, v *gocui.View) error {
	moveViewCursorUp(g, v, 2)
	line, err := getViewLine(g, v)
	POD = getPodNameFromLine(line)
	debug(g, " - Select up in pods view: "+POD)
	return err
}

// View pods: Down
func actionViewPodsDown(g *gocui.Gui, v *gocui.View) error {
	moveViewCursorDown(g, v, false)
	line, err := getViewLine(g, v)
	POD = getPodNameFromLine(line)
	debug(g, " - Select down in pods view: "+POD)
	return err
}

// View pods: Delete
func actionViewPodsDelete(g *gocui.Gui, v *gocui.View) error {
	debug(g, "Delete pod: "+POD)
	err := deletePod(POD)

	go viewPodsRefreshList(g)

	return err
}

// Show views logs
func showViewLogs(g *gocui.Gui, v *gocui.View) error {
	vn := "logs"

	debug(g, "Action: Show view logs")
	g.SetViewOnTop(vn)
	g.SetViewOnTop(vn + "-containers")
	g.SetCurrentView(vn)

	// TODO Enable logs
	switch LOG_MOD {
	case "pod":
		v, err := g.View(vn)
		if err != nil {
			return err
		}
		getPodLogs(POD, v)
	}

	return nil
}

// View pods: Logs
func actionViewPodsLogs(g *gocui.Gui, v *gocui.View) error {
	LOG_MOD = "pod"
	err := showViewLogs(g, v)

	return err
}

// View logs: Hide
func actionViewLogsHide(g *gocui.Gui, v *gocui.View) error {
	g.SetViewOnBottom("logs")
	g.SetViewOnBottom("logs-containers")
	g.SetCurrentView("pods")

	v.Clear()

	debug(g, "Action: Hide view logs)")

	return nil
}
