package storage

import "io"

type Type string
type CloudType = Type

const OSS CloudType = "oss"
const R2 CloudType = "r2"
const AWS CloudType = "aws"
const LOCAL Type = "localfs"

type Storager interface {
	SendFile(pathKey string, file io.Reader, cType string) (string, error)
	SendContent(pathKey string, content []byte) (string, error)
}
