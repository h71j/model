package invite_code

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *Model) GetByCode(ctx context.Context, code string) (*Model, error) {
	// because you should avoid to export something to customers
	coll := m.Context.Handler.Collection(m.Context.Collection)
	filter := bson.D{{Key: "code", Value: code}}

	// 获取数据列表
	// 绑定查询结果
	result := &Model{}

	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
