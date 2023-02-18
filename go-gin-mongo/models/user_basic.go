package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type UserBasic struct {
	Identity  string `bson:"identity"`
	Account   string `bson:"account"`
	Password  string `bson:"password"`
	Nickname  string `bson:"nickname"`
	Sex       int    `bson:"sex"`
	Email     string `bson:"email"`
	Avatar    string `bson:"avatar"`
	CreatedAt int    `bson:"created_at"`
	UpdatedAt int    `bson:"updated_at"`
}

func (UserBasic) CollectionName() string {
	return "user_basic"
}

func GetUserBasicByAccountPassword(account, password string) (*UserBasic, error) {
	ub := new(UserBasic)
	fmt.Println(account, password)
	err := Mongo.Collection(UserBasic{}.CollectionName()).FindOne(context.Background(), bson.D{
		{Key: "account", Value: account}, {Key: "password", Value: password}}).Decode(ub)
	if err != nil {
		return nil, err
	}
	return ub, nil
}

func GetUserBasicByIdentity(identity string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).FindOne(context.Background(), bson.D{
		{Key: "identity", Value: identity}}).Decode(ub)
	return ub, err
}

func GetUserBasicByAccount(account string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).FindOne(context.Background(), bson.D{
		{Key: "account", Value: account}}).Decode(ub)
	return ub, err
}
