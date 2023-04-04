package core

import (
	"context"

	"blogs/app/core/api/internal/svc"
	"blogs/app/core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPhotoAlbumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPhotoAlbumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPhotoAlbumsLogic {
	return &GetPhotoAlbumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPhotoAlbumsLogic) GetPhotoAlbums(req *types.GetPhotoAlbumsRequest) ([]*types.GetPhotoAlbumsResponse, error) {
	res := l.svcCtx.PhotoAlbumDao.GetPhotoAlbums(0, 10)
	albums := make([]*types.GetPhotoAlbumsResponse, 0)
	for _, album := range res {
		albums = append(albums, &types.GetPhotoAlbumsResponse{
			Id:         album.Id,
			AblumCover: album.AlbumCover,
			AlbumDesc:  album.AlbumDesc,
			AlbumName:  album.AlbumName,
		})
	}

	return albums, nil
}
