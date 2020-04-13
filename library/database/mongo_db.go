package database

import (
	"context"
	"fmt"
	"bell/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(conf *config.Config) (*mongo.Client,error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		conf.Mongodb.Username,
		conf.Mongodb.Password,
		conf.Mongodb.Host,
		conf.Mongodb.Port)
	clientOptions := options.Client().ApplyURI(uri)
	client ,err :=mongo.NewClient(clientOptions)
	if err!=nil{
		return nil,err
	}
	err = client.Ping(context.TODO(), nil)
	if err!=nil{
		return nil,err
	}
	return client,nil
}

// 返回mongo_db对应的数据库
func NewDatabaseMongo(conf *config.Config,client *mongo.Client) *mongo.Database {
	return client.Database(conf.Mongodb.Database)
}
