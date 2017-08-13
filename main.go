package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	// Define UI
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	// Configure global UI
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	// Define the layout
	g.SetManagerFunc(uiLayout)

	// Let's go dude!
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// Define the UI
func uiLayout(g *gocui.Gui) error {
	return nil
}
