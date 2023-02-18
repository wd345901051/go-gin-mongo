package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRoom struct {
	UserIdentity    string `bson:"user_identity"`
	RoomIdentity    string `bson:"room_identity"`
	MessageIdentity string `bson:"message_ideneity"`
	RoomType        int    `bson:"room_type"` //[1-独聊，2-公聊]
	CreatedAt       int    `bson:"created_at"`
	UpdatedAt       int    `bson:"updated_at"`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}

func GetUserRoomByUserIdentityRoomIdentity(userIdentity string, roomIdentity string) (*UserRoom, error) {
	ur := new(UserRoom)
	err := Mongo.Collection(UserRoom{}.CollectionName()).FindOne(context.Background(), bson.D{{Key: "user_identity", Value: userIdentity}, {Key: "room_idendity", Value: roomIdentity}}).
		Decode(ur)
	if err != nil {
		return nil, err
	}
	return ur, nil
}

func GetUserRoomByRoomIdentity(roomIdentity string) ([]*UserRoom, error) {
	curosr, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{{Key: "room_identity", Value: roomIdentity}})
	if err != nil {
		return nil, err
	}
	urs := make([]*UserRoom, 0)
	for curosr.Next(context.Background()) {
		ur := new(UserRoom)
		err := curosr.Decode(ur)
		if err != nil {
			return nil, err
		}
		urs = append(urs, ur)
	}
	return urs, nil
}

func JudgeUserIsFriend(userIdentity1, userIdentity2 string) bool {
	// 查询 userIdentity1 单聊
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{{Key: "user_identity", Value: userIdentity1}, {Key: "room_type", Value: 1}})
	if err != nil {
		log.Printf("DB ERROR:%v\n", err)
		return false
	}
	roomIdentitys := make([]*UserRoom, 0)
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			log.Printf("DB ERROR:%v\n", err)
			return false
		}
		roomIdentitys = append(roomIdentitys, ur)
	}

	// 查询 userIdentity2 单聊
	cnt, err := Mongo.Collection(UserRoom{}.CollectionName()).CountDocuments(context.Background(), bson.D{{Key: "user_identity", Value: userIdentity2}, {Key: "room_identity", Value: bson.D{{Key: "$in", Value: roomIdentitys}}}})
	if err != nil {
		log.Printf("DB ERROR:%v\n", err)
		return false
	}
	if cnt > 0 {
		return true
	}
	return false
}
