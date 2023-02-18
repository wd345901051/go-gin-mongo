package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageBasic struct {
	UserIdentity string `bson:"user_identity"`
	RoomIdentity string `bson:"room_identity"`
	Data         string `bson:"data"`
	CreatedAt    int    `bson:"created_at"`
	UpdatedAt    int    `bson:"updated_at"`
}

func (MessageBasic) CollectionName() string {
	return "message_basic"
}

func InsertOneMessageBasic(mb *MessageBasic) error {
	_, err := Mongo.Collection(MessageBasic{}.CollectionName()).InsertOne(context.Background(), mb)
	if err != nil {
		return err
	}
	return nil
}

func GetMessageListByRoomIdentity(RoomIdentity string, limit, skip int64) ([]*MessageBasic, error) {
	data := make([]*MessageBasic, 0)
	cursor, err := Mongo.Collection(MessageBasic{}.CollectionName()).Find(context.Background(), bson.M{"room_ideneity": RoomIdentity}, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  bson.D{{Key: "created_at", Value: -1}},
	})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		mb := new(MessageBasic)
		err := cursor.Decode(mb)
		if err != nil {
			return nil, err
		}
		data = append(data, mb)
	}
	return data, nil
}
