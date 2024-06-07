package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func (m *MongoDB) Connect(uri, dbName string) {
	var err error
	for {
		clientOptions := options.Client().ApplyURI(uri)
		m.Client, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Println("Failed to connect to MongoDB, retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}

		err = m.Client.Ping(context.TODO(), nil)
		if err != nil {
			log.Println("Failed to ping MongoDB, retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	m.Database = m.Client.Database(dbName)
	log.Println("Successfully connected to MongoDB")
}
