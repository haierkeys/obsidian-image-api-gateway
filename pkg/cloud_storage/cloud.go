package cloud_storage

import "io"

type CloudStorage interface {
	SendFile(pathKey string, file io.Reader, cType string) (string, error)
	SendContent(pathKey string, content []byte) (string, error)
}
