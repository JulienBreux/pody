package main

import (
	"fmt"

	"github.com/willf/pad"
)

var HELP = `
Pody controls the Pods on your Kubernetes cluster.
Find more information at https://github.com/JulienBreux/pody.
The following options can be passed to any command:

  --help        for more information about Pody, hmm, here btw!
  --version     information about the version
  --frequency   refreshing frequency in seconds (default: 5)
  --kubeconfig  absolute path to the kubeconfig file
`

const APP = "Pody"
const AUTHOR = "@JulienBreux"

// Get full banner of version
func versionFull() string {
	return fmt.Sprintf("%s %s - By %s", APP, version, AUTHOR)
}

// Get only banner (used in title bar view)
func versionBanner() string {
	return fmt.Sprintf(" %s %s", APP, version)
}

// Get only author (used in title bar view)
func versionAuthor() string {
	return fmt.Sprintf("By %s  ", AUTHOR)
}

// Prepare version title (used in title bar view)
func versionTitle(width int) string {
	return "⣿" + versionBanner() + pad.Left(versionAuthor(), width-len(versionBanner()), " ")
}
