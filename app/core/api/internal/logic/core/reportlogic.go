package core

import (
	"blogs/common/constants"
	"blogs/common/utils"
	"blogs/common/utils/useragent"
	"blogs/lib/cache"
	"context"
	"github.com/WillVi/ipgo"
	"net"
	"net/http"
	"strings"

	"blogs/app/core/api/internal/svc"
	"blogs/app/core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportLogic {
	return &ReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReportLogic) Report(r *http.Request) (resp *types.ReportResponse, err error) {
	ipAddress := getIpAddress(r)
	userAgent := r.UserAgent()
	ua := useragent.Parse(userAgent)
	browser := ua.Name
	os := ua.OS

	uuid := ipAddress + browser + os
	md5 := utils.MD5(uuid)

	if !cache.SIsMember(constants.UniqueVisitorKey, md5) {
		ipgo.New("./resource/ipdb/ip2region.db")
		search, _ := ipgo.BtreeSearch(ipAddress)
		ipSource := strings.ReplaceAll(strings.ReplaceAll(search.Region, "|0", ""), "0|", "")
		if len(ipSource) > 0 {
			province := getProvince(ipSource)
			cache.HIncr(constants.VisitorsAreaDetailKey, province)
		} else {
			cache.HIncr(constants.VisitorsAreaDetailKey, constants.UnknownArea)
		}
		cache.IncrCount(constants.BlogsViewCountKey)
		cache.SAdd(constants.UniqueVisitorKey, md5)
	}

	return
}

func getProvince(ipSource string) string {
	if len(ipSource) == 0 {
		return constants.UnknownArea
	}

	strs := strings.Split(ipSource, "|")
	if len(strs) > 1 && strings.HasSuffix(strs[1], constants.AreaSuffix) {
		return strs[1][:strings.Index(strs[1], constants.AreaSuffix)]
	}
	return constants.DefaultArea
}

func getIpAddress(r *http.Request) string {
	ipAddress := r.Header.Get("X-Real-IP")

	if len(ipAddress) == 0 || ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = r.Header.Get("x-forwarded-for")
	}

	if len(ipAddress) == 0 || ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = r.Header.Get("Proxy-Client-IP")
	}

	if len(ipAddress) == 0 || ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = r.Header.Get("WL-Proxy-Client-IP")
	}

	if len(ipAddress) == 0 || ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = r.Header.Get("HTTP_CLIENT_IP")
	}

	if len(ipAddress) == 0 || ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = r.Header.Get("HTTP_X_FORWARDED_FOR")
	}

	if len(ipAddress) == 0 || ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = r.RemoteAddr
		if "127.0.0.1" == ipAddress || "0:0:0:0:0:0:0:1" == ipAddress {
			ipAddress = getLocalIpV4()
		}
	}
	return ipAddress
}

func getLocalIpV4() string {
	inters, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, inter := range inters {
		// 判断网卡是否开启，过滤本地环回接口
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			// 获取网卡下所有的地址
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					//判断是否存在IPV4 IP 如果没有过滤
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}