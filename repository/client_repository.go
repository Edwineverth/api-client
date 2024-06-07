package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

	"api-client/dao"
)

type ClientRepository struct {
	ClientDAO *dao.ClientDAO
}

func (r *ClientRepository) GetAllClients(ctx context.Context) ([]map[string]interface{}, error) {
	return r.ClientDAO.GetAllClients(ctx)
}

func (r *ClientRepository) GetClientByID(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.ClientDAO.GetClientByID(ctx, id)
}

func (r *ClientRepository) CreateClient(ctx context.Context, client map[string]interface{}) (*mongo.InsertOneResult, error) {
	return r.ClientDAO.CreateClient(ctx, client)
}

func (r *ClientRepository) UpdateClient(ctx context.Context, id string, update map[string]interface{}) (*mongo.UpdateResult, error) {
	return r.ClientDAO.UpdateClient(ctx, id, update)
}

func (r *ClientRepository) DeleteClient(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	return r.ClientDAO.DeleteClient(ctx, id)
}
