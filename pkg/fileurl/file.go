package fileurl

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type FileType int

const TypeImage FileType = iota + 1

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()

}

// GetFileName 获取文件路径
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	// fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

// GetFileExt 获取文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetDatePath 获取日期保存路径
func GetDatePath() string {
	getYearMonth := time.Now().Format("200601")
	getDay := time.Now().Format("02")
	return getYearMonth + "/" + getDay + "/"
}

// IsContainExt 判断文件后缀是否在允许范围内
func IsContainExt(t FileType, name string, uploadAllowExts []string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range uploadAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

// IsFileSizeAllowed 文件大小是否被允许
func IsFileSizeAllowed(t FileType, f multipart.File, uploadMaxSize int) bool {
	content, _ := io.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= uploadMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// IsPermission 检查文件权限
func IsPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// IsExist 判断所给路径是否不存在
func IsExist(dst string) bool {
	_, err := os.Stat(dst) // os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// CreatePath 创建路径
func CreatePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

// GetExePath 获取当前执行文件的路径
func GetExePath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}

// PathSuffixCheckAdd 检查路径后缀，如果没有则添加
func PathSuffixCheckAdd(path string, suffix string) string {
	if !strings.HasSuffix(path, suffix) {
		path = path + "/"
	}
	return path
}

// IsAbsPath 判断是否为绝对路径
func IsAbsPath(path string) bool {
	if runtime.GOOS == "windows" {
		// Windows系统
		if filepath.VolumeName(path) != "" {
			return true // 如果有盘符，则为绝对路径
		}
		return filepath.IsAbs(path) // 检查是否是绝对路径
	}
	// UNIX/Linux系统
	return filepath.IsAbs(path)
}

// GetAbsPath 获取绝对路径
func GetAbsPath(path string, root string) (string, error) {
	if root != "" {
		root = PathSuffixCheckAdd(root, "/")
	}
	realPath := root + path
	// 如果本身就是绝对路径 就直接返回
	if !IsAbsPath(realPath) {
		pwdDir, _ := os.Getwd()
		realPath = PathSuffixCheckAdd(pwdDir, "/") + path
	}
	if IsExist(realPath) {
		return realPath, nil
	} else {
		return "", errors.New("file not exists")
	}
}
