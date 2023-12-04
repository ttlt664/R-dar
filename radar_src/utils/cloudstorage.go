package utils

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"radar_src/config"
)

var OSSClient *oss.Client

func Init() error {
	// 从配置文件中读取云存储配置
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// 创建 OSS 客户端
	client, err := oss.New(cfg.OSS.Endpoint, cfg.OSS.AccessKey, cfg.OSS.SecretKey)
	if err != nil {
		return err
	}

	OSSClient = client
	return nil
}

// GetBucket 获取 OSS 存储桶
func GetBucket() (*oss.Bucket, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	// 获取存储桶
	bucket, err := OSSClient.Bucket(cfg.OSS.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
