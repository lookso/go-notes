package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
	if err != nil {
		panic(err)
	}
	db := client.Database("cron")
	coll := db.Collection("log")
	count, err := coll.Find(ctx, bson.M{"jobName": "抓取QQ音乐字典"}).Count()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}
