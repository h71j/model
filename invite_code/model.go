package invite_code

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "member_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_info"
	// 这个需要用户根据具体业务完成设定
	modelName = "invite_code"
)

// Model 邀请码表
type Model struct {
	// 模型继承
	model.Model `json:"basic" bson:"basic"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// 邀请码
	Code string `json:"code" bson:"code"`
	// 有效时间
	ExpireAt string `json:"expire_at" bson:"expire_at"`

	// 邀请者获取奖励数量
	RewardCnt int `json:"reward_cnt" bson:"reward_cnt"`

	// 邀请用户列表
	Invitees []*Invitee `json:"invite_users" bson:"invite_users"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}

// Invitee 被邀请的用户
type Invitee struct {
	// 用户id
	AccountID string `json:"account_id"  bson:"account_id"`

	// 创建时间
	CreatedAt string `json:"created_at"  bson:"created_at"`
}
