package global

import (
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/fileurl"
)

var (
	// 程序执行目录
	ROOT string
	Name string = "obsidian image-api gateway"
)

func init() {

	filename := fileurl.GetExePath()
	ROOT = filename + "/"

}
