package global

import (
    "github.com/haierkeys/obsidian-image-api-gateway/pkg/fsutil"
)

var (
    // 程序执行目录
    ROOT string
    Name string = "obsidian image-api gateway"
)

func init() {

    filename := fsutil.GetExePath()
    ROOT = filename + "/"

}
