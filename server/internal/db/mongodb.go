package db

// import (
// 	"context"
// 	"path"
// 	"strings"

// 	"go.mongodb.org/mongo-driver/v2/mongo"
// 	"go.mongodb.org/mongo-driver/v2/mongo/options"
// 	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
// )

// // implements Database
// type MongoDb struct {
// 	uri    string
// 	dbName string
// 	client *mongo.Client
// }

// func NewMongoDb(uri string) Database {
// 	dbName := path.Base(strings.Split(uri, "?")[0])
// 	return &MongoDb{
// 		uri:    uri,
// 		dbName: dbName,
// 	}
// }

// func (m *MongoDb) Connect() error {
// 	client, err := mongo.Connect(options.Client().ApplyURI(m.uri))
// 	if err != nil {
// 		return err
// 	}
// 	m.client = client
// 	return nil
// }
// func (m *MongoDb) Disconnect() error {
// 	if m.client == nil {
// 		return NotInitializedErr("MongoDb")
// 	}
// 	m.client.Disconnect(context.Background())
// 	return nil
// }

// func (m *MongoDb) Health() error {
// 	if m.client == nil {
// 		return NotInitializedErr("MongoDb")

// 	}
// 	return m.client.Ping(context.Background(), readpref.Primary())
// }

// func (m *MongoDb) Collection(name string) *mongo.Collection {
// 	return m.client.Database(m.dbName).Collection(name)
// }
