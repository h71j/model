package invite

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
	modelName = "invite"
)

// Model 模型
type Model struct {
	// 模型继承
	model.Model `json:"basic" bson:"basic"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// 邀请码
	InviteID string `json:"invite_id" bson:"invite_id"`
	// 账号
	AccountID string `json:"account_id" bson:"account_id"`
	// 当前邀请任务的完成数量
	InviteNum int `json:"invite_num" bson:"invite_num"`
	// 被邀请的列表
	Inviter []string `json:"inviter" bson:"inviter"`
	// 邀请完成后成为朋友
	Friends []string `json:"friends" bson:"friends"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
