// operation.go

package webdav

import (
	"fmt"
	"os"
)

// SendFile 将本地文件上传到 WebDAV 服务器。
func (w *WebDAV) SendFile(localPath, remotePath string) error {
	content, err := os.ReadFile(localPath)
	if err != nil {
		return fmt.Errorf("打开本地文件失败: %v", err)
	}

	err = w.Client.Write(remotePath, content, os.ModePerm)

	if err != nil {
		return fmt.Errorf("上传文件失败: %v", err)
	}

	return nil
}

// SendContent 将二进制内容上传到 WebDAV 服务器。
func (w *WebDAV) SendContent(remotePath string, content []byte) error {

	err := w.Client.Write(remotePath, content, os.ModePerm)

	if err != nil {
		return fmt.Errorf("上传文件失败: %v", err)
	}

	return nil
}

// // DownloadFile 从 WebDAV 服务器下载文件到本地。
// func (w *WebDAV) DownloadFile(remotePath, localPath string) error {
// 	err := w.Client.DownloadFile(remotePath, localPath)
// 	if err != nil {
// 		return fmt.Errorf("下载文件失败: %v", err)
// 	}

// 	return nil
// }

// // DeleteFile 从 WebDAV 服务器删除文件。
// func (w *WebDAV) DeleteFile(remotePath string) error {
// 	err := w.Client.Remove(remotePath)
// 	if err != nil {
// 		return fmt.Errorf("删除文件失败: %v", err)
// 	}

// 	return nil
// }

// // MkDir 在 WebDAV 服务器上创建目录。
// func (w *WebDAV) MkDir(remotePath string) error {
// 	err := w.Client.Mkdir(remotePath)
// 	if err != nil {
// 		if !gowebdav.IsErrExist(err) {
// 			return fmt.Errorf("创建目录失败: %v", err)
// 		}
// 		// 如果目录已存在，忽略错误
// 		log.Printf("目录 %s 已存在，忽略错误", remotePath)
// 	}

// 	return nil
// }

// // ListFiles 列出 WebDAV 服务器上的文件和目录。
// func (w *WebDAV) ListFiles(remotePath string) ([]string, error) {
// 	files, err := w.Client.ReadDir(remotePath)
// 	if err != nil {
// 		return nil, fmt.Errorf("列出文件失败: %v", err)
// 	}

// 	var fileNames []string
// 	for _, file := range files {
// 		fileNames = append(fileNames, file.Name())
// 	}

// 	return fileNames, nil
// }
