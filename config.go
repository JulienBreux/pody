package main

import (
	"flag"
	"os"
	"path/filepath"
)

type config struct {
	homeDir    string
	kubeConfig *string
	frequency  int
	setup      bool
}

// Check if configuration is initialized
func (c *config) Initialized() bool {
	return c != nil && c.setup
}

var CONFIG config

// Get configuration
func getConfig() config {
	if CONFIG.Initialized() {
		return CONFIG
	}

	c := config{}

	// Home directory
	c.homeDir = os.Getenv("USERPROFILE")
	if h := os.Getenv("HOME"); h != "" {
		c.homeDir = h
	}

	// Kubernetes configuration
	if h := c.homeDir; h != "" {
		c.kubeConfig = flag.String("kubeconfig", filepath.Join(h, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		c.kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// Refreshing frequency
	flag.IntVar(&c.frequency, "frequency", 3, "refreshing frequency in seconds (default: 5)")

	flag.Parse()

	c.setup = true
	CONFIG = c

	return c
}
