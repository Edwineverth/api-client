package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientDAO struct {
	Collection *mongo.Collection
}

func (dao *ClientDAO) GetAllClients(ctx context.Context) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	cursor, err := dao.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result map[string]interface{}
		cursor.Decode(&result)
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (dao *ClientDAO) GetClientByID(ctx context.Context, id string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := dao.Collection.FindOne(ctx, bson.M{"clientId": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (dao *ClientDAO) CreateClient(ctx context.Context, client map[string]interface{}) (*mongo.InsertOneResult, error) {
	return dao.Collection.InsertOne(ctx, client)
}

func (dao *ClientDAO) UpdateClient(ctx context.Context, id string, update map[string]interface{}) (*mongo.UpdateResult, error) {
	filter := bson.M{"clientId": id}
	updateBson := bson.M{"$set": update}
	return dao.Collection.UpdateOne(ctx, filter, updateBson)
}

func (dao *ClientDAO) DeleteClient(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	return dao.Collection.DeleteOne(ctx, bson.M{"clientId": id})
}
