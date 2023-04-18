package logic

import (
	"blogs/models/db"
	"context"
	"sort"

	"blogs/app/admin/api/internal/svc"
	"blogs/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMenusLogic {
	return &UserMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserMenusLogic) UserMenus(userId int64) ([]types.GetMenusResponse, error) {
	menus := l.svcCtx.MenuDao.GetMenusByUserId(userId)
	level1Dirs := make([]types.GetMenusResponse, 0)
	for _, menu := range menus {
		if menu.ParentId == 0 {
			level1Dirs = append(level1Dirs, types.GetMenusResponse{
				Name:      menu.Name,
				Path:      menu.Path,
				Component: menu.Component,
				Icon:      menu.Icon,
				Hidden:    menu.IsHidden == 1,
				OrderNum:  menu.OrderNum,
				Children:  getMenuById(menu.Id, menus),
			})
		}
	}
	sort.Slice(level1Dirs, func(i, j int) bool {
		return level1Dirs[i].OrderNum < level1Dirs[j].OrderNum
	})
	return level1Dirs, nil
}

func getMenuById(menuId int64, menus []db.Menu) []types.GetMenusResponse {
	resp := make([]types.GetMenusResponse, 0)
	for _, menu := range menus {
		if menu.ParentId == menuId {
			resp = append(resp, types.GetMenusResponse{
				Name:      menu.Name,
				Path:      menu.Path,
				Component: menu.Component,
				Icon:      menu.Icon,
				Hidden:    menu.IsHidden == 0,
				OrderNum:  menu.OrderNum,
				Children:  getMenuById(menu.Id, menus),
			})
		}
	}
	sort.Slice(resp, func(i, j int) bool {
		return resp[i].OrderNum < resp[j].OrderNum
	})
	return resp
}
