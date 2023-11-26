package menu

import (
	"github.com/h71j/model/goods"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Model) GetGoods(goodsIds []string) ([]*goods.Model, error) {

	results := make([]*goods.Model, 0)
	coll := m.Context.Handler.Collection(m.Context.Collection)
	objsIds := make([]*primitive.ObjectID, 0)
	for _, i := range goodsIds {
		objID, _ := primitive.ObjectIDFromHex(i)
		objsIds = append(objsIds, &objID)
	}

	filter := bson.M{"$in": objsIds}

	// 获取数据列表
	cursor, err := coll.Find(m.Context.Context, filter)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if err = cursor.All(m.Context.Context, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// GetGoodsByMenuID 后续改为批量查询
func (m *Model) GetGoodsByMenuID(menuID string) ([]*goods.Model, error) {

	err := m.GetOne(m, menuID)
	if err != nil {
		return nil, err
	}

	goodsIds := make([]string, 0)
	for _, i := range m.GoodsList {
		goodsIds = append(goodsIds, i.GoodsId)
	}

	results := make([]*goods.Model, 0)
	coll := m.Context.Handler.Collection(m.Context.Collection)
	objsIds := make([]*primitive.ObjectID, 0)
	for _, i := range m.Goods {
		objID, _ := primitive.ObjectIDFromHex(i)
		objsIds = append(objsIds, &objID)
	}

	filter := bson.M{"$in": objsIds}

	// 获取数据列表
	cursor, err := coll.Find(m.Context.Context, filter)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if err = cursor.All(m.Context.Context, &results); err != nil {
		return nil, err
	}
	return results, nil
}
