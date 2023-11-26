package menu

import (
	"github.com/h71j/model/goods"
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "product_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "menu"
)

// Model 商品信息
type Model struct {
	// 模型继承
	model.Model `json:"-" bson:"-"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Name            string `json:"name"`
	IsShowBackstage int    `json:"is_show_backstage" bson:"is_show_backstage"`
	Sort            int    `json:"sort"`
	GoodsType       int    `json:"goods_type"  bson:"goods_type"`
	IsSell          bool   `json:"is_sell" bson:"is_sell"`
	Icon            string `json:"icon"`
	// 即将废弃 goods_list
	Goods     []string      `json:"goods_list" bson:"goods_list"`
	GoodsList []GoodsConfig `json:"goods_list_v2" bson:"goods_list_v2"`
	// GoodsDisplay 货物展示 不存储
	GoodsDisplay []*goods.Model `json:"goods_display" bson:"-"`
	// 发布到的门店
	Stores []string `json:"stores" bson:"stores"`
	// 更新方式
	UpdateType int `json:"update_type" bson:"update_type"`
}

type GoodsConfig struct {
	// Sales 销量
	Sales int `json:"sales"`
	// Stock 库存
	Stock int `json:"stock"`
	// GoodsId
	GoodsId string `json:"goods_id"  bson:"goods_id"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
