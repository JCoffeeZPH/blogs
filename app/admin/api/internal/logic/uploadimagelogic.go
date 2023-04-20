package logic

import (
	"blogs/common/config"
	"blogs/common/constants"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"

	"blogs/app/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageLogic {
	return &UploadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImageLogic) UploadImage(r *http.Request) (string, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return "", err
	}
	files := r.MultipartForm.File["file"]
	filename := files[0].Filename
	extName := filename[strings.LastIndex(filename, ".")+1:]

	logx.Infof("UploadImageHandler file details, filename: %s, extName: %s, filesize: %d", filename, extName, files[0].Size)
	file, err := files[0].Open()
	if err != nil {
		logx.Errorf("open file failed, err: %+v", err)
		return "", err
	}

	fileBytes, _ := io.ReadAll(file)
	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	md5Value := hex.EncodeToString(hash.Sum(nil))

	filePath := fmt.Sprintf("%s/%s.%s", constants.OSSArticleImagesPath, md5Value, extName)
	return fmt.Sprintf("%s/%s", config.GetCfg().OSS.BucketUrl, filePath), l.svcCtx.OSService.UploadFile(filePath, fileBytes)
}
