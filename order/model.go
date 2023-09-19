package order

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "order_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_flow"
	// 这个需要用户根据具体业务完成设定
	modelName = "command"
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
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Status           int                `json:"status"`
	CompletedTime    string             `json:"completed_time" bson:"completed_time"`
	MultiStore       string             `json:"multi_store" bson:"multi_store"`
	ProductionedTime string             `json:"productioned_time" bson:"productioned_time"`
	TypeCate         int                `json:"typeCate" bson:"typeCate"`
	SendStatus       int                `json:"send_status" bson:"send_status"`
	QueueIndex       int64              `json:"queue_index" bson:"queue_index"`

	User     UserInfo       `json:"user"  bson:"user"`
	Order    OrderInfo      `json:"order"  bson:"order"`
	Pay      PayInfo        `json:"pay"  bson:"pay"`
	Store    StoreInfo      `json:"store"  bson:"store"`
	Discount []DiscountInfo `json:"discount"  bson:"discount"`
	Goods    []GoodInfo     `json:"goods" bson:"goods"`
	TakeOut  TakeOut        `json:"take_out" bson:"take_out"`
}

type UserInfo struct {
	Mobile   string `json:"mobile"`
	UserName string `json:"user_name"`
}

type OrderInfo struct {
	UpdatedAt   int    `json:"updated_at" bson:"updated_at"`
	GoodsNum    int    `json:"goods_num" bson:"goods_num"`
	Status      int    `json:"status" bson:"status"`
	CompletedAt int    `json:"completed_at" bson:"completed_at"`
	CreatedAt   int    `json:"created_at" bson:"created_at"`
	SendedTime  int    `json:"sended_time" bson:"sended_time"`
	Remark      string `json:"remark"`
	Postscript  string `json:"postscript" bson:"postscript"`
	SortNum     string `json:"sort_num" bson:"sort_num"`
	OrderNo     string `json:"order_no" bson:"order_no"`
	StatusText  string `json:"status_text" bson:"status_text"`
}

type PayInfo struct {
	CouponAmount string `json:"coupon_amount"  bson:"coupon_amount"`
	PayedAt      int    `json:"payed_at" bson:"payed_at"`
	TotalAmount  string `json:"total_amount" bson:"total_amount"`
	CouponName   string `json:"coupon_name" bson:"coupon_name"`
	ReceiveAt    int    `json:"receive_at" bson:"receive_at"`
	PayMode      string `json:"pay_mode" bson:"pay_mode"`
	Amount       string `json:"amount"`
	PayUserName  string `json:"pay_user_name" bson:"pay_user_name"`
}

type DiscountInfo struct {
	Summary string `json:"summary"`
	Amount  string `json:"amount"`
	Method  string `json:"method"`
	OrderNo string `json:"order_no"`
	Name    string `json:"name"`
	DataId  string `json:"data_id"`
}

type StoreInfo struct {
	Address   string `json:"address"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Mobile    string `json:"mobile"`
	Name      string `json:"name"`
	StoreID   string `json:"store_id" bson:"store_id"`
}
type GoodInfo struct {
	Number       int    `json:"number"`
	OriginAmount string `json:"origin_amount"  bson:"origin_amount"`
	Price        string `json:"price"`
	Unit         string `json:"unit"`
	Property     string `json:"property"`
	Image        string `json:"image"`
	Amount       string `json:"amount"`
	Name         string `json:"name"`
}

// TakeOut 外卖订单信息
type TakeOut struct {
	ReturnCode string `json:"return_code"` // 外卖单号
	PriceToken string `json:"price_token"` // 外卖价格token
	Price      string `json:"price"`       // 跑腿费
	Receiver   string `json:"receiver"`    // 收件人
	RecPhone   string `json:"rec_phone"`   // 手机号码
	Remark     string `json:"remark"`      // 备注
	Address    string `json:"address"`     // 地址
	Lng        string `json:"lng"`         // 经度
	Lat        string `json:"lat"`         // 维度
	Distance   string `json:"distance"`    // 距离
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
