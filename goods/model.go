package goods

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
	modelName = "goods"
)

// Model 商品信息
type Model struct {
	// 模型继承
	model.Model `json:"-" bson:"-"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID      int64              `json:"product_id" bson:"product_id"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
	GoodsMealsInfo []interface{}      `json:"goods_meals_info,omitempty" bson:"goods_meals_info"`
	IsAdd          int                `json:"is_add,omitempty" bson:"is_add"`
	UseSpec        bool               `json:"use_spec" bson:"use_spec"`
	Entity         []EntityInfo       `json:"entity"`
	StallCode      string             `json:"stall_code" bson:"stall_code"`
	Specs          []SpecsInfo        `json:"specs"`
	IsFollowSuit   int                `json:"is_follow_suit,omitempty"`
	IsLabel        int                `json:"is_label"`
	// 销售属性
	SellTimeStatus  int     `json:"sell_time_status" bson:"sell_time_status"`
	IsSell          bool    `json:"is_sell" bson:"is_sell"`
	PackCost        string  `json:"pack_cost" bson:"pack_cost"`
	UnitType        int     `json:"unit_type" bson:"unit_type"`
	Sort            int     `json:"sort"`
	Price           float64 `json:"price"`
	Unit            string  `json:"unit"`
	MembershipPrice int     `json:"membership_price" bson:"membership_price"`

	// 发布属性
	PublishClient []string `json:"publish_client" bson:"publish_client"`

	// 商品本身属性
	Name          string         `json:"name"`
	Type          int            `json:"type"`
	GoodsType     int            `json:"goods_type" bson:"goods_type"`
	Content       string         `json:"content"`
	UseProperty   int            `json:"use_property" bson:"use_property"`
	IsUseProperty bool           `json:"is_use_property" bson:"is_use_property"`
	CategoryID    string         `json:"category_id" bson:"category_id"`
	Property      []PropertyInfo `json:"property"`

	// 统计数据 & 限制
	Sales     int `json:"sales"`
	Stock     int `json:"stock"`
	MinBuyNum int `json:"min_buy_num" bson:"min_buy_num"`

	// 展示
	Images   string   `json:"images"`
	CoverImg string   `json:"cover_img" bson:"cover_img"`
	ImageArr []string `json:"imageArr"`
}

type SpecsInfo struct {
	Values []struct {
		Id    int         `json:"id"`
		Image interface{} `json:"image"`
		Value string      `json:"value"`
	} `json:"values"`
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type PropertyInfo struct {
	IsOpenCheckbox bool              `json:"is_open_checkbox"`
	Id             int               `json:"id"`
	Values         []PropertySetting `json:"values"`
	Name           string            `json:"name"`
	Desc           *string           `json:"desc,omitempty"`
}

type PropertySetting struct {
	IsDefault int     `json:"is_default,omitempty"`
	Id        int     `json:"id"`
	Code      string  `json:"code"`
	Value     string  `json:"value"`
	Price     float64 `json:"price"`
}

type EntityInfo struct {
	SpecId          string      `json:"spec_id"`
	TradeMark       string      `json:"trade_mark"`
	Id              string      `json:"id"`
	Stock           string      `json:"stock"`
	SpecText        interface{} `json:"spec_text"`
	Spec            interface{} `json:"spec"`
	Image           string      `json:"image"`
	Num             int         `json:"num"`
	Price           float64     `json:"price"`
	MembershipPrice int         `json:"membership_price"  bson:"membership_price"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
