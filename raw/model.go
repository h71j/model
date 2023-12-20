package raw

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "product_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_item"
	// 这个需要用户根据具体业务完成设定
	modelName = "raw"
)

// Model 商品信息
type Model struct {
	// 模型继承
	model.Model `json:"-" bson:"-"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID        int64              `json:"product_id" bson:"product_id"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at"`
	RawId            int                `json:"raw_id"`
	RawName          string             `json:"raw_name"`
	Mnemonic         string             `json:"mnemonic"`
	Unit             string             `json:"unit"`
	RawSubId         int                `json:"raw_sub_id"`
	RawSubName       string             `json:"raw_sub_name"`
	Category         string             `json:"category"`
	StdUnit          string             `json:"std_unit"`
	StdQrCode        string             `json:"std_qr_code"`
	OrderUnit        string             `json:"order_unit"`
	OrderQrCode      string             `json:"order_qr_code"`
	ConversionRate   float64            `json:"conversion_rate"`
	CostUnit         string             `json:"cost_unit"`
	CostConversion   float64            `json:"cost_conversion"`
	PurchasingUnit   string             `json:"purchasing_unit"`
	Barcode          string             `json:"barcode"`
	Status           string             `json:"status"`
	IsTheBatchManage string             `json:"is_the_batch_manage"`
	IsTheShelfLife   string             `json:"is_the_shelf_life"`
	ShelfLifeDays    int                `json:"shelf_life_days"`
	RemindBeforeDays int                `json:"remind_before_days"`
	InStockExpire    int                `json:"in_stock_expire"`
	OutStockExpire   int                `json:"out_stock_expire"`
	Rate             string             `json:"rate"`
	StoreCondition   string             `json:"store_condition"`
	Remark           string             `json:"remark"`
	Tag              []string           `json:"tag"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
