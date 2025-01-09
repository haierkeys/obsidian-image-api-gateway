package storage

import (
	"io"

	"github.com/haierkeys/obsidian-image-api-gateway/pkg/code"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/storage/aliyun_oss"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/storage/aws_s3"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/storage/cloudflare_r2"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/storage/local_fs"
)

type Type = string
type CloudType = Type

const OSS CloudType = "oss"
const R2 CloudType = "r2"
const S3 CloudType = "s3"
const LOCAL Type = "localfs"

var StorageTypeMap = map[Type]bool{
	OSS:   true,
	R2:    true,
	S3:    true,
	LOCAL: true,
}

var CloudStorageTypeMap = map[Type]bool{
	OSS: true,
	R2:  true,
	S3:  true,
}

type Storager interface {
	SendFile(pathKey string, file io.Reader, cType string) (string, error)
	SendContent(pathKey string, content []byte) (string, error)
}

var Instance map[Type]Storager

func NewClient(cType Type, config map[string]any) (Storager, error) {

	if cType == LOCAL {
		return local_fs.NewClient(config)
	} else if cType == OSS {
		return aliyun_oss.NewClient(config)
	} else if cType == R2 {
		return cloudflare_r2.NewClient(config)
	} else if cType == S3 {
		return aws_s3.NewClient(config)
	}
	return nil, code.ErrorInvalidStorageType
}
