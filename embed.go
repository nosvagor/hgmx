package hgmx

import (
	"embed"
)

//go:embed .version
var version string

//go:embed static/** components/**
var StaticFS embed.FS

func Version() string {
	return version
}
