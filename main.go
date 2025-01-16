package main

import (
	"embed"

	"github.com/haierkeys/obsidian-image-api-gateway/cmd"
)

//go:embed frontend
var efs embed.FS

func main() {
	cmd.Execute(efs)
}
