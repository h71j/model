package attendance

import (
	"context"
	"errors"
	"github.com/h71j/model/member"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func GetToday() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02")
}

// Integrals 积分记录
func (m *Model) Integrals(ctx context.Context, openID string, pointNum int) (*mongo.UpdateResult, error) {

	// 查询会员信息
	filterForMember := bson.D{{Key: "external.open_id", Value: openID}}
	memberModel := &member.Model{}
	memberColl := m.Context.Handler.Collection(memberModel.Context.Collection)
	err := memberColl.FindOne(ctx, filterForMember).Decode(memberModel)
	if err != nil {
		return nil, err
	}

	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := m.Context.Handler.Collection(m.Context.Collection)
	filter := bson.D{{Key: "meta.account_id", Value: memberModel.Meta.AccountID}, {Key: "date", Value: GetToday()}}
	attendanceModel := &Model{}
	// attendanceModel 查询签到记录
	err = coll.FindOne(ctx, filter).Decode(&attendanceModel)

	if err == mongo.ErrNoDocuments {
		// 未签到过，可以进行签到
		m.Meta.AccountID = memberModel.Meta.AccountID
		m.PointNum = pointNum
		m.IsAttendance = true
		m.Date = GetToday()

		_, err2 := m.Create(m)
		if err2 != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	if attendanceModel.IsAttendance {
		return nil, errors.New("already attendance")
	}

	// 获取数据列表
	// 绑定查询结果

	// 充值金额
	memberModel.Assets.PointNum = memberModel.Assets.PointNum + pointNum

	updateResult, err := memberColl.UpdateOne(ctx, filterForMember, bson.D{{Key: "$set", Value: memberModel}})
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}
