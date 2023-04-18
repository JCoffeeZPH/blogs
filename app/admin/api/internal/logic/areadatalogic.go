package logic

import (
	"blogs/common/constants"
	"blogs/lib/cache"
	"context"
	"strconv"

	"blogs/app/admin/api/internal/svc"
	"blogs/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AreaDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAreaDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AreaDataLogic {
	return &AreaDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AreaDataLogic) AreaData(req *types.GetAreaDataRequest) ([]*types.GetAreaDataResponse, error) {
	var data map[string]string
	if req.Type == constants.DataAreaVisitorType {
		data = cache.HGetAll(constants.UserAreaDetailKey)
	} else {
		data = cache.HGetAll(constants.VisitorsAreaDetailKey)
	}

	resp := make([]*types.GetAreaDataResponse, 0)
	for key, value := range data {
		num, _ := strconv.ParseInt(value, 10, 64)
		resp = append(resp, &types.GetAreaDataResponse{
			Name:  key,
			Value: num,
		})
	}

	return resp, nil
}
