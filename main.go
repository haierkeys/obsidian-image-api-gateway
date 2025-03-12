package main

import (
	"embed"

	"github.com/haierkeys/obsidian-image-api-gateway/cmd"
)

//go:embed frontend
var efs embed.FS

//go:embed config/config.yaml
var c string

func main() {
	cmd.Execute(efs, c)
}
