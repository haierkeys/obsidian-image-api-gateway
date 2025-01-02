package global

import (
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/path"
)

var (
	// 程序执行目录
	ROOT string
	Name string = "obsidian image-api gateway"
)

func init() {

	filename := path.GetExePath()
	ROOT = filename + "/"

}
