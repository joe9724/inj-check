package model

type CheckModel struct{
	ID int64 `json:"id"`
	UserID int64 `json:"user_id"`
	CreateTime string `json:"create_time"`
	Status int64 `json:"status"`
}

//启动页model
type StartUpModel struct{
	// download Url
	DownloadURL string `json:"downloadUrl,omitempty"`
	// force
	Force string `json:"force,omitempty"`

	// msg
	Msg string `json:"msg,omitempty"`

	// number
	Number float64 `json:"number,omitempty"`

	Cover string `json:"cover"`

	CoverX string `json:"coverx"`

	ShareUrl string `json:"shareUrl"`

}

//首页model细分
//Banner
type Banner struct {
	// id
	ID int64 `json:"id"`

	Title string `json:"title"`
	// cover
	Cover string `json:"cover"`
	// type
	Type int64 `json:"type"`
	// 跳转类型(0=圈子 1=帖子 2=用户 3=活动 4=新闻 5=web)

	TargetID int64 `json:"target_id"`

	WebUrl string `json:"webUrl"`

	EventID string `json:"event_id"`

}
type Icon struct {
	// icon
	Icon string `json:"icon"`

	// id
	ID int64 `json:"id"`

	// name
	Title string `json:"title"`

	// order
	Order int64 `json:"order"`

	// 跳转类型(0=圈子 1=帖子 2=用户 3=活动 4=新闻 5=web)
	Type int64 `json:"type"`

	// web Url
	WebURL string `json:"web_url"`

	TargetID int64 `json:"target_id"`
}

type CityToutiao struct {
	Cover string `json:"cover"`
	// icon
	Icon string `json:"icon"`

	// id
	ID int64 `json:"id"`

	// name
	Title string `json:"title"`

	SubTitle string `json:"sub_title"`

	// order
	Order int64 `json:"order"`

	// 跳转类型(0=圈子 1=帖子 2=用户 3=活动 4=新闻 5=web)
	Type int64 `json:"type"`

	// web Url
	WebURL string `json:"web_url"`

	TargetID int64 `json:"target_id"`

	Style string `json:"style"`

	Status int64 `json:"status"`

	EventID string `json:"event_id"`
}
type ZoneItem struct {

	// 圈子简介
	Brief string `json:"brief"`

	// 创建日期
	CreateAt string `json:"create_at"`

	// 是否加入该圈子
	Joind bool `json:"joind,omitempty"`

	// 圈子等级
	Level int64 `json:"level"`

	// 圈子的缩略图
	// Required: true
	Logo *string `json:"logo"`

	// 圈子的成员个数
	MemberCount int64 `json:"member_count"`

	// 圈子名称
	// Required: true
	Name *string `json:"name"`

	// 圈子状态
	Status int64 `json:"status"`

	// 圈子的标签
	Tag string `json:"tag"`

	// 圈子唯一标识符
	// Required: true
	ZoneID int64 `json:"zone_id"`

	Cover string `json:"cover"`
}
type Activity struct{
	EventThumb string `json:"event_thumb"`
	EventTitle string `json:"event_title"`
	EventSummary string `json:"event_summary"`
	EventID string `json:"event_id"`
	Tag string `json:"tag"`
}
type News struct{
	ID int64 `json:"id"`
	Thumb string `json:"thumb"`
	Title string `json:"title"`
	Source string `json:"source"`
	CreateTime string `json:"create_time"`
	WebUrl string `json:"web_url"`
}
type Home struct{
	Banners []Banner `json:"banners"`
	Icons []Icon `json:"icons"`
	// CityTT CityToutiao `json:"cityToutiao"`
	Zones []ZoneItem `json:"zones"`
	Activities []Activity `json:"activities"`
	//News []News `json:"news"`
	CityTopList []CityToutiao `json:"cityToplist"`
	CityBanner CityToutiao `json:"cityBanner"`
}
