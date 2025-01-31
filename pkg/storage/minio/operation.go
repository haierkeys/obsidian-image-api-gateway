package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/haierkeys/obsidian-image-api-gateway/pkg/fileurl"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/pkg/errors"
)

func (p *MinIO) GetBucket(bucketName string) string {

	// Get bucket
	if len(bucketName) <= 0 {
		bucketName = p.Config.BucketName
	}

	return bucketName
}

// UploadByFile 上传文件
func (p *MinIO) SendFile(fileKey string, file io.Reader, itype string) (string, error) {

	ctx := context.Background()
	bucket := p.GetBucket("")

	fileKey = fileurl.PathSuffixCheckAdd(p.Config.CustomPath, "/") + fileKey

	//  k, _ := h.Open()

	_, err := p.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileKey),
		Body:        file,
		ContentType: aws.String(itype),
	})

	if err != nil {
		return "", errors.Wrap(err, "minio")
	}

	return fileurl.PathSuffixCheckAdd(p.Config.BucketName, "/") + fileKey, nil
}

func (p *MinIO) SendContent(fileKey string, content []byte) (string, error) {

	ctx := context.Background()
	bucket := p.GetBucket("")

	fileKey = fileurl.PathSuffixCheckAdd(p.Config.CustomPath, "/") + fileKey

	input := &s3.PutObjectInput{
		Bucket:            aws.String(bucket),
		Key:               aws.String(fileKey),
		Body:              bytes.NewReader(content),
		ChecksumAlgorithm: types.ChecksumAlgorithmSha256,
	}
	output, err := p.S3Manager.Upload(ctx, input)
	if err != nil {
		var noBucket *types.NoSuchBucket
		if errors.As(err, &noBucket) {
			fmt.Printf("Bucket %s does not exist.\n", bucket)
			err = noBucket
		}
	} else {
		err := s3.NewObjectExistsWaiter(p.S3Client).Wait(ctx, &s3.HeadObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(fileKey),
		}, time.Minute)
		if err != nil {
			fmt.Printf("Failed attempt to wait for object %s to exist in %s.\n", fileKey, bucket)
		} else {
			_ = *output.Key
		}
	}

	return fileurl.PathSuffixCheckAdd(p.Config.BucketName, "/") + fileKey, errors.Wrap(err, "minio")
}
