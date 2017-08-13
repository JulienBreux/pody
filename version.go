package main

import (
	"fmt"

	"github.com/willf/pad"
)

const APP = "Pody"
const AUTHOR = "@JulienBreux"
const VERSION = "0.1.0"

func versionFull() string {
	return fmt.Sprintf("%s %s - By %s", APP, VERSION, AUTHOR)
}

func versionBanner() string {
	return fmt.Sprintf("  %s %s", APP, VERSION)
}

func versionAuthor() string {
	return fmt.Sprintf("By %s  ", AUTHOR)
}

func versionTitle(width int) string {
	return versionBanner() + pad.Left(versionAuthor(), width-len(versionBanner()), " ")
}
