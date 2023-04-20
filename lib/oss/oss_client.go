package oss

import (
	"blogs/common/config"
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/core/logx"
)

type OSService interface {
	UploadFile(path string, fileBytes []byte) error
}

type OSServiceImpl struct {
	ossClient *oss.Client
}

func NewOSService() OSService {
	cfg := config.GetCfg()

	ossClient, err := oss.New(cfg.OSS.Endpoint, cfg.OSS.AccessKeyID, cfg.OSS.AccessKeySecret)
	if err != nil {
		panic(fmt.Errorf("init oss client failed, err: %+v", err))
	}
	return &OSServiceImpl{
		ossClient: ossClient,
	}
}

func (impl *OSServiceImpl) UploadFile(path string, fileBytes []byte) error {
	bucket, err := impl.ossClient.Bucket(config.GetCfg().OSS.BucketName)
	if err != nil {
		logx.Errorf("get bucket failed when upload to oss, err: %+v", err)
		return err
	}
	exist, err := bucket.IsObjectExist(path)
	if err != nil {
		return err
	}
	if exist {
		logx.Infof("%s already exist", path)
		return nil
	}
	return bucket.PutObject(path, bytes.NewReader(fileBytes))
}
