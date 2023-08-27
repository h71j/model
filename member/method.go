package member

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetByExternal 详情
// GetByExternal	GET http://my.api.url/posts/123
func (m *Model) GetByExternal(ctx context.Context, id string) (*Model, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := m.Context.Handler.Collection(m.Context.Collection)
	filter := bson.D{{Key: "external.open_id", Value: id}}

	// 获取数据列表
	// 绑定查询结果
	result := &Model{}

	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TopUp 充值
func (m *Model) TopUp(ctx context.Context, id string, amount float64) (*mongo.UpdateResult, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := m.Context.Handler.Collection(m.Context.Collection)
	filter := bson.D{{Key: "external.open_id", Value: id}}

	// 获取数据列表
	// 绑定查询结果
	oldResult := &Model{}
	err := coll.FindOne(ctx, filter).Decode(&oldResult)
	if err != nil {
		return nil, err
	}

	// 充值金额
	oldResult.Assets.Balance = amount

	updateResult, err := coll.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: oldResult}})
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

// Integrals 积分记录
func (m *Model) Integrals(ctx context.Context, id string, pointNum int) (*mongo.UpdateResult, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := m.Context.Handler.Collection(m.Context.Collection)
	filter := bson.D{{Key: "external.open_id", Value: id}}

	// 获取数据列表
	// 绑定查询结果
	oldResult := &Model{}
	err := coll.FindOne(ctx, filter).Decode(&oldResult)
	if err != nil {
		return nil, err
	}

	// 充值金额
	oldResult.Assets.PointNum = pointNum

	updateResult, err := coll.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: oldResult}})
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}
