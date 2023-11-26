package material

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
	collectionNameSuffix = "_info"
	// 这个需要用户根据具体业务完成设定
	modelName = "material"
)

// Model 商品信息
type Model struct {
	// 模型继承
	model.Model `json:"-" bson:"-"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID int64              `json:"product_id" bson:"product_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	// 单位（斤/盒/箱/桶
	UnitType int `json:"unit_type" bson:"unit_type"`
	// 单价
	UnitPrice float64 `json:"unit_price"  bson:"unit_price"`
	// 名称（油、面、米）
	Name string `json:"name"  bson:"name"`
	// Remark 备注
	Remark string `json:"remark"  bson:"remark"`
	// 类型
	Type int `json:"type"`
	// Origin 产地
	Origin string `json:"origin" bson:"origin"`
	// Stock 库存
	Stock int `json:"stock"`
	// StockName 库存统计单位
	StockName string `json:"stock_name" bson:"stock_name"`
	// 图片
	Icon string `json:"icon"`
	// 供应商列表
	Supplier []string `json:"supplier"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
