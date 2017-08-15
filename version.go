package main

import (
	"fmt"

	"github.com/willf/pad"
)

const APP = "Pody"
const AUTHOR = "@JulienBreux"
const VERSION = "0.1.0"

// Get full banner of version
func versionFull() string {
	return fmt.Sprintf("%s %s - By %s", APP, VERSION, AUTHOR)
}

// Get only banner (used in title bar view)
func versionBanner() string {
	return fmt.Sprintf("  %s %s", APP, VERSION)
}

// Get only author (used in title bar view)
func versionAuthor() string {
	return fmt.Sprintf("By %s  ", AUTHOR)
}

// Prepare version title (used in title bar view)
func versionTitle(width int) string {
	return versionBanner() + pad.Left(versionAuthor(), width-len(versionBanner()), " ")
}
