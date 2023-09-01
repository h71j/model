package order

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetListDesc 获取列表
// getList	GET http://my.api.url/posts?sort=["title","ASC"]&range=[0, 24]&filter={"title":"bar"}
func (m *Model) GetListDesc(filter interface{}, d interface{}) (int64, error) {
	coll := m.Context.Handler.Collection(m.Context.Collection)
	// 声明需要返回的列表
	// 获取总数（含过滤规则）
	totalCounter, err := coll.CountDocuments(context.TODO(), filter)
	if err == mongo.ErrNoDocuments {
		return 0, err
	}

	opt := options.Find()
	// 排序方式
	opt.SetSort(bson.M{"order.created_at": -1})

	// 获取数据列表
	cursor, err := coll.Find(m.Context.Context, filter, opt)
	if err == mongo.ErrNoDocuments {
		return totalCounter, err
	}

	if err != nil {
		return totalCounter, err
	}

	if err = cursor.All(context.TODO(), d); err != nil {
		return totalCounter, err
	}
	return totalCounter, nil
}
