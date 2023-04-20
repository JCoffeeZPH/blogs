package constants

type Namespace string

const (
	DevRPCNameSpaceId     Namespace = "d7532de5-68a7-4e3d-abf1-a4344761a064"
	LiveRPCNameSpaceId    Namespace = "9b47b39b-c2ba-42a7-8a91-80458ab88c42"
	StagingRPCNameSpaceId Namespace = "9ef8d701-4a16-4f83-a7a7-0c00b6ced270"
	TestRPCNameSpaceId    Namespace = "abb7dca5-f180-4c75-9585-a31475732b74"
	UatRPCNameSpaceId     Namespace = "bd2867d9-b8e0-4c63-ae51-29b2f80f04cd"

	BlogRPCServerName = "blog.rpc"

	UnknownArea = "未知"
	DefaultArea = "广东"
	AreaSuffix  = "省"

	DataAreaVisitorType = 1
	DataAreaUserType    = 0

	OSSArticleImagesPath = "articles"
)

var namespaceMap = map[string]Namespace{
	"dev":     DevRPCNameSpaceId,
	"live":    LiveRPCNameSpaceId,
	"staging": StagingRPCNameSpaceId,
	"test":    TestRPCNameSpaceId,
	"uat":     UatRPCNameSpaceId,
}

func GetNamespace(env string) Namespace {
	return namespaceMap[env]
}

func (ns Namespace) ToString() string {
	return string(ns)
}
