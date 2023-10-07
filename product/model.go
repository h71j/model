package product

import (
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
	modelName = "goods"
)

// LbsInfo 地址
type LbsInfo struct {
	Address   string  `json:"address" bson:"address"`
	Longitude float64 `json:"longitude"  bson:"longitude"`
	Latitude  float64 `json:"latitude"  bson:"latitude"`
	AreaName  string  `json:"areaName"  bson:"areaName"`
}

// Model 商品信息
type Model struct {
	// 模型继承
	model.Model `json:"-" bson:"-"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Name            string      `json:"name"`
	IsShowBackstage int         `json:"is_show_backstage" bson:"is_show_backstage"`
	Sort            int         `json:"sort"`
	GoodsType       int         `json:"goods_type"  bson:"goods_type"`
	IsSell          bool        `json:"is_sell" bson:"is_sell"`
	Icon            string      `json:"icon"`
	Goods           []GoodsInfo `json:"goods_list" bson:"goods_list"`
	// 发布到的门店
	Stores []string `json:"stores" bson:"stores"`
	// 更新方式
	UpdateType int `json:"update_type" bson:"update_type"`
}

type GoodsInfo struct {
	SellTimeStatus  int            `json:"sell_time_status" bson:"sell_time_status"`
	Id              int            `json:"id"`
	IsSell          bool           `json:"is_sell" bson:"is_sell"`
	PackCost        string         `json:"pack_cost" bson:"pack_cost"`
	Sales           int            `json:"sales"`
	GoodsType       int            `json:"goods_type" bson:"goods_type"`
	CoverImg        string         `json:"cover_img" bson:"cover_img"`
	Property        []PropertyInfo `json:"property"`
	GoodsMealsInfo  []interface{}  `json:"goods_meals_info,omitempty" bson:"goods_meals_info"`
	IsAdd           int            `json:"is_add,omitempty" bson:"is_add"`
	UseSpec         bool           `json:"use_spec" bson:"use_spec"`
	Entity          []EntityInfo   `json:"entity"`
	StallCode       string         `json:"stall_code" bson:"stall_code"`
	Sort            int            `json:"sort"`
	Price           float64        `json:"price"`
	Unit            string         `json:"unit"`
	ImageArr        []string       `json:"imageArr"`
	MembershipPrice int            `json:"membership_price" bson:"membership_price"`
	UseProperty     int            `json:"use_property" bson:"use_property"`
	IsUseProperty   bool           `json:"is_use_property" bson:"is_use_property"`
	UnitType        int            `json:"unit_type" bson:"unit_type"`
	MinBuyNum       int            `json:"min_buy_num" bson:"min_buy_num"`
	Specs           []struct {
		Values []struct {
			Id    int         `json:"id"`
			Image interface{} `json:"image"`
			Value string      `json:"value"`
		} `json:"values"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"specs"`
	Content      string `json:"content"`
	IsFollowSuit int    `json:"is_follow_suit,omitempty"`
	Stock        string `json:"stock"`
	Type         int    `json:"type"`
	IsLabel      int    `json:"is_label"`
	Name         string `json:"name"`
	Images       string `json:"images"`
}

type PropertyInfo struct {
	IsOpenCheckbox bool `json:"is_open_checkbox"`
	Id             int  `json:"id"`
	Values         []struct {
		IsDefault int    `json:"is_default,omitempty"`
		Id        int    `json:"id"`
		Code      string `json:"code"`
		Value     string `json:"value"`
	} `json:"values"`
	Name string  `json:"name"`
	Desc *string `json:"desc,omitempty"`
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
