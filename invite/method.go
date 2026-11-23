package invite

import (
	"context"
	"fmt"

	"github.com/bwmarrin/snowflake"
)

// InitInvite 邀请列表
func (m *Model) InitInvite(ctx context.Context, accountID string) error {
	// because you should avoid to export something to customers
	m.AccountID = accountID
	m.InviteID = newInviteID() // 暂时使用这个id
	_, err := m.Create(m)
	if err != nil {
		return err
	}
	return nil
}

// InviteJoin 邀请列表
func (m *Model) InviteJoin(accountID string) error {
	m.AccountID = accountID
	m.InviteID = newInviteID()
	_, err := m.Create(m)
	if err != nil {
		return err
	}

	//filter := bson.D{{Key: "account_id", Value: accountID}}

	// 获取数据列表
	// 绑定查询结果
	//oldResult := &Model{}
	//err := coll.FindOne(ctx, filter).Decode(&oldResult)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// 充值金额
	//oldResult.Assets.Balance = amount
	//
	//updateResult, err := coll.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: oldResult}})
	//if err != nil {
	//	return nil, err
	//}
	return nil
}

// inviteNode 雪花 ID 节点（与原 r2day/base/util.GetBrandId 相同：节点号 1，前缀 "b"）
var inviteNode, inviteNodeErr = snowflake.NewNode(1)

// newInviteID 生成邀请 ID，格式与原 r2day/base/util.GetBrandId 保持一致（"b" + 雪花 ID）
func newInviteID() string {
	if inviteNodeErr != nil {
		return "-"
	}
	return fmt.Sprintf("b%d", inviteNode.Generate())
}
