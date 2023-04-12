package logic

import (
	"blogs/app/admin/api/internal/svc"
	"blogs/app/admin/api/internal/types"
	"blogs/common/errorx"
	"blogs/common/utils"
	"blogs/lib/cache"
	sess "blogs/models/session"
	"context"
	"github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest, jwtToken string) (*types.LoginResponse, error) {
	var session *sess.Session
	if jwtToken != "" {
		s := utils.ParseSession(jwtToken)
		if s != nil && s.Username == req.Username {
			session = s
		}
	}

	if session == nil {
		user := l.svcCtx.UserAuthDao.GetUserInfo(req.Username, req.Password)
		if user == nil {
			logx.Errorf("VerifyToken not found user")
			panic(errorx.NotFound)
		}

		sessionId := uuid.NewV1().String()
		secret := uuid.NewV1().String()
		s := &sess.Session{
			SessionId:     sessionId,
			UserId:        user.UserInfoId,
			Username:      user.Username,
			Secret:        secret,
			Nickname:      user.Nickname,
			Email:         user.Email,
			Avatar:        user.Avatar,
			Role:          user.RoleId,
			LastLoginTime: user.LastLoginTime,
			IpAddress:     user.IpAddress,
			IpSource:      user.IpSource,
			IsSubscribe:   user.IsSubscribe,
			LoginType:     user.LoginType,
			Website:       user.Website,
		}
		session = s
	}

	go cache.SaveSessionIntoRedis(session)

	return createLoginResp(session), nil
}

func createLoginResp(s *sess.Session) *types.LoginResponse {
	return &types.LoginResponse{
		Id:            s.UserId,
		Avatar:        s.Avatar,
		Email:         s.Email,
		Intro:         s.Nickname,
		IpAddress:     s.IpAddress,
		IpSource:      s.IpSource,
		IsSubscribe:   s.IsSubscribe,
		LastLoginTime: utils.TimeFormat(s.LastLoginTime),
		LoginType:     s.LoginType,
		Nickname:      s.Nickname,
		Username:      s.Username,
		Website:       s.Website,
		Token:         utils.GenerateJwtToken(s.SessionId, s.Secret),
	}

}
