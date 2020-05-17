package config

import (
	config "github.com/pcelvng/go-config"
)

// Options determine where the server looks for
// different files
type Options struct {
	StaticDir   string `flag:"static-dir"`
	TemplateDir string `flag:"template-dir"`
	HTTPPort    string `flag:"http-port"`
}

// Site holds various information regarding the site
type Site struct {
	Name    string  `flag:"site-name"`
	Options Options `flag:"opts"`
}

// Load site configuration from environment and cli
func Load(site *Site) error {
	return config.Load(site)
}
