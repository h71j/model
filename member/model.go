package member

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "affiliate_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_info"
	// 这个需要用户根据具体业务完成设定
	modelName = "account"
)

// LbsInfo 地址
type LbsInfo struct {
	Address   string  `json:"address" bson:"address"`
	Longitude float64 `json:"longitude"  bson:"longitude"`
	Latitude  float64 `json:"latitude"  bson:"latitude"`
	AreaName  string  `json:"areaName"  bson:"areaName"`
}

// Model 门店信息
type Model struct {
	// 模型继承
	model.Model `json:"basic" bson:"basic"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// 第三方账号信息
	External ExternalInfo `json:"external" bson:"external"`
	// 手机号
	Phone string `json:"phone"`
	// 更多信息
	// 账号名称
	Nickname string     `json:"nickname"`
	Avatar   string     `json:"avatar"`
	Member   MemberInfo `json:"member" bson:"member"`
	Assets   AssetsInfo `json:"assets" bson:"assets"`
}

type ExternalInfo struct {
	OpenID string `json:"open_id"  bson:"open_id"`
}

type MemberInfo struct {
	Level    int    `json:"level"`
	Gender   int    `json:"gender"`
	Birthday string `json:"birthday"`
	CardName string `json:"card_name" bson:"card_name"`
	CardUrl  string `json:"cardUrl" bson:"cardUrl"`
}

type AssetsInfo struct {
	PointNum     int     `json:"pointNum" bson:"pointNum"`
	CouponNum    int     `json:"couponNum" bson:"couponNum"`
	Balance      float64 `json:"balance" bson:"balance"`
	GiftBalance  int     `json:"giftBalance" bson:"giftBalance"`
	CurrentValue int     `json:"currentValue" bson:"currentValue"`
	NeedValue    int     `json:"needValue" bson:"needValue"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
