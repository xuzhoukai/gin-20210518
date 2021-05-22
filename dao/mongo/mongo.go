package dao_mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBDataBase *mongo.Database

func init() {
	uri := "mongodb://gin01:gin01@139.224.0.98:27017/gin01"
	dbName := "gin01"
	timeOut := 5 * time.Second
	var maxConnect uint64 = 100

	_, err := ConnectToDB(uri, dbName, timeOut, maxConnect)
	if err != nil {
		log.Fatalf("connect to mongodb err:%v", err)
	}
}

func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, err
	}
	// 主从模式 随机主从选择服务器
	mod := readpref.Nearest()
	// 主机模式 从master读取 默认模式
	mod = readpref.Primary()
	ops := options.Database()
	ops.SetReadPreference(mod)

	MongoDBDataBase = client.Database(name, ops)
	return client.Database(name), nil
}
