package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
)

var DEBUG_DISPLAYED bool = false
var LOG_MOD string = "pod"
var NAMESPACE string = "default"
var POD string = ""

// Configure globale keys
var keys []Key = []Key{
	Key{"", gocui.KeyCtrlC, actionGlobalQuit},
	Key{"", gocui.KeyCtrlD, actionGlobalToggleDebug},
	Key{"pods", gocui.KeyArrowUp, actionViewPodsUp},
	Key{"pods", gocui.KeyArrowDown, actionViewPodsDown},
	Key{"pods", 'd', actionViewPodsDelete},
	Key{"pods", 'l', actionViewPodsLogs},
	Key{"logs", 'l', actionViewLogsHide},
}

// Main or not main, that's the question^^
func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(uiLayout)

	if err := uiKey(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// Define the UI layout
func uiLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	viewDebug(g, maxX, maxY)
	viewLogs(g, maxX, maxY)
	viewOverlay(g, maxX, maxY)
	viewTitle(g, maxX, maxY)
	viewPods(g, maxX, maxY)

	return nil
}

// Useful to debug Pody (display with CTRL+D)
func debug(g *gocui.Gui, output interface{}) {
	v, err := g.View("debug")
	if err == nil {
		t := time.Now()
		tf := t.Format("2006-01-02 15:04:05")
		output = tf + " > " + output.(string)
		fmt.Fprintln(v, output)
	}
}

// Move view cursor to the bottom
func moveViewCursorDown(g *gocui.Gui, v *gocui.View, allowEmpty bool) error {
	cx, cy := v.Cursor()
	ox, oy := v.Origin()
	nextLine, err := getNextViewLine(g, v)
	if err != nil {
		return err
	}
	if !allowEmpty && nextLine == "" {
		return nil
	}
	if err := v.SetCursor(cx, cy+1); err != nil {
		if err := v.SetOrigin(ox, oy+1); err != nil {
			return err
		}
	}
	return nil
}

// Move view cursor to the top
func moveViewCursorUp(g *gocui.Gui, v *gocui.View, dY int) error {
	ox, oy := v.Origin()
	cx, cy := v.Cursor()
	if cy > dY {
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

// Get view line (relative to the cursor)
func getViewLine(g *gocui.Gui, v *gocui.View) (string, error) {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	return l, err
}

// Get the next view line (relative to the cursor)
func getNextViewLine(g *gocui.Gui, v *gocui.View) (string, error) {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy + 1); err != nil {
		l = ""
	}

	return l, err
}

// Set view cursor to line
func setViewCursorToLine(g *gocui.Gui, v *gocui.View, lines []string, selLine string) error {
	ox, _ := v.Origin()
	cx, _ := v.Cursor()
	for y, line := range lines {
		if line == selLine {
			if err := v.SetCursor(ox, y); err != nil {
				if err := v.SetOrigin(cx, y); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Get pod name form line
func getPodNameFromLine(line string) string {
	if line == "" {
		return ""
	}

	i := strings.Index(line, " ")
	if i == -1 {
		return line
	}

	return line[0:i]
}
