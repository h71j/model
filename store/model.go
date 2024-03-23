package store

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "mini_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_manage"
	// 这个需要用户根据具体业务完成设定
	modelName = "store"
)

// BusinessOperationModel 是商业运营模式的枚举类型
type BusinessOperationModel int

const (
	// CompanyOwned 直营模式
	CompanyOwned BusinessOperationModel = iota
	// Franchise 加盟模式
	Franchise
	// Licensing 特许经营模式
	Licensing
	// Agency 代理商模式
	Agency
	// Distribution 分销模式
	Distribution
	// Partnership 合作伙伴关系模式
	Partnership
	// JointVenture 联营模式
	JointVenture
)

// OpenScale 表示门店规模的枚举
type OpenScale int

const (
	// Small 小型门店
	Small OpenScale = iota
	// Medium 中型门店
	Medium
	// Large 大型门店
	Large
)

// LbsInfo 地址
type LbsInfo struct {
	Address   string  `json:"address" bson:"address"`
	Longitude float64 `json:"longitude"  bson:"longitude"`
	Latitude  float64 `json:"latitude"  bson:"latitude"`
	AreaName  string  `json:"areaName"  bson:"areaName"`
}

// ShopTimeSetting 营业时间
type ShopTimeSetting struct {
	BeginTime string `json:"begin_time" bson:"begin_time"`
	EndTime   string `json:"end_time" bson:"end_time"`
}

// TypeInfo 门店类型
type TypeInfo struct {
	// 商业运营模式的枚举类型
	BOM BusinessOperationModel `json:"bom" bson:"bom"`
	// 门店规模
	OS OpenScale `json:"os" bson:"os"`
}

// Model 门店信息
type Model struct {
	// 模型继承
	model.Model `json:"-" bson:"-"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// 分类/ 亦或则是分组等
	Category string `json:"category" bson:"category"`
	// BestMatch  最匹配的门店
	BestMatch bool `json:"best_match"`
	// 门店状态
	Status int `json:"status"  bson:"status"`
	// StoreID 门店id
	StoreID string `json:"storeId" bson:"storeId"`
	// Name 门店名称
	Name string `json:"name" bson:"name"`
	// Distance 根据用户当前位置计算出的结果
	Distance float64 `json:"distance" bson:"-"`
	// 门店营业时间(客户端展示）
	ShopTime string `json:"shopTime" bson:"shopTime"`
	// ShopTimeSetting 门店营业时间
	ShopTimeInfo ShopTimeSetting `json:"shop_time_info" bson:"shop_time_info"`
	// 门店电话
	CallNumber string `json:"callNumber" bson:"callNumber"`
	// 地址
	Lbs LbsInfo `json:"lbs" bson:"lbs"`
	// 当前门店拥有的产品
	Product []string `json:"product" bson:"product"`
	// 门店的餐桌二维码
	TableQrCodes []Qrcode     `json:"table_qrcodes"  bson:"table_qrcodes"`
	MenuList     []MenuConfig `json:"menu_list" bson:"menu_list"`
	// Printer 打印机配置
	Printer PrinterConf `json:"printer_conf" bson:"printer_conf"`
	// 财务主体
	Finance FinanceConfig `json:"finance" bson:"finance"`
	// 门店类型属性设置
	ST TypeInfo `json:"st" bson:"st"`
}

type FinanceConfig struct {
	// WxPay 微信支付
	WxPay string `json:"wx_pay"  bson:"wx_pay"`
	// 银行支付
	BankPay string `json:"bank_pay" bson:"bank_pay"`
}

type MenuConfig struct {
	// Sales 销量
	Sales int `json:"sales"`
	// Stock 库存
	Stock int `json:"stock"`
	// MenuID
	MenuID string `json:"menu_id"  bson:"menu_id"`
}

type PrinterConf struct {
	// 前台打印机
	Front Printer `json:"front"  bson:"front"`
	// 后厨打印机
	Kitchen Printer `json:"kitchen"  bson:"kitchen"`
}

type Printer struct {
	Sn      string `json:"sn"  bson:"sn"`
	User    string `json:"user"  bson:"user"`
	UserKey string `json:"user_key"  bson:"user_key"`
	Debug   string `json:"debug"  bson:"debug"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
