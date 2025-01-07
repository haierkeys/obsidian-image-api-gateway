package fsutil

import (
	"errors"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	// fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

type FileType int

const TypeImage FileType = iota + 1

func FileToMultipart(file *os.File) (multipart.File, *multipart.FileHeader, error) {

	// 将 *os.File 对象转换为 multipart.File 类型
	fileInfo, _ := file.Stat()
	return file, &multipart.FileHeader{
		Filename: fileInfo.Name(),
		Size:     fileInfo.Size(),
		// ModTime:  fileInfo.ModTime(),
		// 如果还需要其他属性，可以根据实际情况进行设置
	}, nil
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePreDirPath() string {

	getYearMonth := time.Now().Format("200601")
	getDay := time.Now().Format("02")
	return getYearMonth + "/" + getDay + "/"
}

func UrlEscape(fileKey string) string {

	if strings.Contains(fileKey, "/") {

		i := strings.LastIndex(fileKey, "/")

		fileKey = fileKey[:i+1] + url.PathEscape(fileKey[i+1:])
	} else {
		fileKey = url.PathEscape(fileKey)
	}

	return fileKey
}

func CheckContainExt(t FileType, name string, uploadAllowExts []string) bool {
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

func CheckMaxSize(t FileType, f multipart.File, uploadMaxSize int) bool {
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

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)

	return os.IsPermission(err)
}

func CheckPath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CreatePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

func GetExePath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}
func Exists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// PathSuffixCheckAdd 检查路径后缀，如果没有则添加
func PathSuffixCheckAdd(path string, suffix string) string {
	if !strings.HasSuffix(path, suffix) {

		path = path + "/"
	}
	return path
}

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

func GetAbsPath(path string, root string) (string, error) {

	// fsutil = strings.TrimPrefix(fsutil, "/")

	if root != "" {
		root = PathSuffixCheckAdd(root, "/")
	}

	realPath := root + path

	// 如果本身就是绝对路径 就直接返回
	if IsAbsPath(realPath) {

	} else {
		pwdDir, _ := os.Getwd()
		realPath = PathSuffixCheckAdd(pwdDir, "/") + path
	}
	if Exists(realPath) {
		return realPath, nil
	} else {
		return "", errors.New("fsutil not exists")
	}
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()

}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
