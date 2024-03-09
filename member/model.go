package member

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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
	// 个人简介
	Member Profile `json:"member" bson:"member"`
	// 资产管理
	Assets AssetsInfo `json:"assets" bson:"assets"`
	// 安全
	Safe Safe `json:"safe" bson:"safe"`
}

type ExternalInfo struct {
	OpenID string `json:"open_id"  bson:"open_id"`
	// 注册来源 （微信/支付宝/抖音)
	From string `json:"from"  bson:"from"`
}

type Profile struct {
	// 手机号
	Phone string `json:"phone"`
	// 账号名称
	Nickname string `json:"nickname"`
	// 真实姓名
	Name string `json:"name"`
	// 头像
	Avatar string `json:"avatar"`
	// 等级
	Level int `json:"level"`
	// 性别
	Gender int `json:"gender"`
	// 生日
	Birthday string `json:"birthday"`
	// 卡
	CardName string `json:"card_name" bson:"card_name"`
	// 卡地址
	CardUrl string `json:"card_url" bson:"card_url"`
}

type Safe struct {
	// 最近一次登陆
	LastLoginTime time.Time `json:"last_login_time" bson:"last_login_time"`
	// 最近一次访问来源
	LastLoginFrom time.Time `json:"last_login_from" bson:"last_login_from"`
	// 最近一次访问ip
	LastLoginIp time.Time `json:"last_login_ip" bson:"last_login_ip"`
	// 最近一次下单
	LastPlace time.Time `json:"last_place" bson:"last_place"`
}

type AssetsInfo struct {
	// 点数
	PointNum int `json:"pointNum" bson:"pointNum"`
	// 积分
	CouponNum int `json:"couponNum" bson:"couponNum"`
	// 余额
	Balance float64 `json:"balance" bson:"balance"`
	// 礼物
	GiftBalance int `json:"giftBalance" bson:"giftBalance"`
	// 当前邀请人完成值
	CurrentValue int `json:"currentValue" bson:"currentValue"`
	// 还需要完成值
	NeedValue int `json:"needValue" bson:"needValue"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
