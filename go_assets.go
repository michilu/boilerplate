package main

import (
	"github.com/jessevdk/go-assets"
)

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{}, map[string]*assets.File{}, "")
