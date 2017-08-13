package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

// App
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
	return nil
}

// Global action: Quit
func quitAction(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
