package global

import (
	"github.com/haierkeys/custom-image-gateway/pkg/fileurl"
)

var (
	// 程序执行目录
	ROOT string
	Name string = "Obsidian Image API Gateway"
)

func init() {

	filename := fileurl.GetExePath()
	ROOT = filename + "/"

}
