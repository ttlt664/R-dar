package utils

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"radar_src/config"
)

var OSSClient *oss.Client

func Init() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	client, err := oss.New(cfg.OSS.Endpoint, cfg.OSS.AccessKey, cfg.OSS.SecretKey)
	if err != nil {
		return err
	}

	OSSClient = client
	return nil
}

func GetBucket() (*oss.Bucket, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	bucket, err := OSSClient.Bucket(cfg.OSS.BucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
