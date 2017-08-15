package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jroimartin/gocui"
)

var NAMESPACE string = ""

// Configure globale keys
var keys []Key = []Key{
	Key{"", gocui.KeyCtrlC, actionGlobalQuit},
	Key{"", gocui.KeyCtrlD, actionGlobalToggleDebug},
	Key{"pods", gocui.KeyArrowUp, actionViewPodsUp},
	Key{"pods", gocui.KeyArrowDown, actionViewPodsDown},
	Key{"pods", gocui.KeyEnter, actionViewPodsSelect},
}

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

func uiLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	viewDebug(g, maxX, maxY)
	viewOverlay(g, maxX, maxY)
	viewTitle(g, maxX, maxY)
	viewPods(g, maxX, maxY)

	return nil
}

// Useful to debug application (display with CTRL+D)
func debug(g *gocui.Gui, output interface{}) {
	v, err := g.View("debug")
	if err == nil {
		t := time.Now()
		tf := t.Format("2006-01-02 15:04:05")
		output = tf + " > " + output.(string)
		fmt.Fprintln(v, output)
	}
}
