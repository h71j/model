package product

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Model) GetByStoreID(id string) ([]*Model, error) {

	results := make([]*Model, 0)
	coll := m.Meta.Handler.Collection(m.Meta.Collection)
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "meta.merchant_id", Value: objID}}

	// 获取数据列表
	cursor, err := coll.Find(m.Meta.Context, filter)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if err = cursor.All(m.Meta.Context, &results); err != nil {
		return nil, err
	}
	return results, nil
}
