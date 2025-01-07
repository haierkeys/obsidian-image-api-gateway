package cloudflare_r2

import (
	"context"
	"fmt"

	"github.com/haierkeys/obsidian-image-api-gateway/global"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/pkg/errors"
)

func NewClient() (*s3.Client, error) {
	// New client

	var accountId = global.Config.CloudfluR2.AccountId
	var accessKeyId = global.Config.CloudfluR2.AccessKeyID
	var accessKeySecret = global.Config.CloudfluR2.AccessKeySecret

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {

		return nil, errors.Wrap(err, "cloudflare_r2")
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId))
	})

	if err != nil {
		return nil, errors.Wrap(err, "cloudflare_r2")
	}
	return client, nil
}
