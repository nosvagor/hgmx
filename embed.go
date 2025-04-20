package hgmx

import (
	"embed"
)

//go:embed .version
var version string

//go:embed static/**
var StaticFS embed.FS

func Version() string {
	return version
}
