package main

import (
	"embed"

	"github.com/haierkeys/custom-image-gateway/cmd"
)

//go:embed frontend
var efs embed.FS

//go:embed config/config.yaml
var c string

func main() {
	cmd.Execute(efs, c)
}
