package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type UserInfo struct {
	Name   string `bson:"name"`
	Age    int    `bson:"age"`
	Weight int    `bson:"weight"`
}

var userInfo = UserInfo{
	Name:   "xm",
	Age:    7,
	Weight: 40,
}

func main() {
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://127.0.0.1:27017", Database: "cron", Coll: "log"})
	defer func() {
		if err = cli.Close(ctx); err != nil {
			panic(err)
		}
	}()
	count, err := cli.Find(ctx, bson.M{"name": userInfo.Name}).Count()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)

	//cli.CreateOneIndex(context.Background(), options.IndexModel{Key: []string{"name"}, Unique: true})
	//cli.CreateIndexes(context.Background(), []options.IndexModel{{Key: []string{"id2", "id3"}}})

	// insert one document
	_, err = cli.InsertOne(ctx, userInfo)
	if err != nil {
		fmt.Println("result err", err)
	}

	// multiple insert
	var userInfos = []UserInfo{
		UserInfo{Name: "a1", Age: 6, Weight: 20},
		UserInfo{Name: "b2", Age: 6, Weight: 25},
		UserInfo{Name: "c3", Age: 6, Weight: 30},
		UserInfo{Name: "d4", Age: 6, Weight: 35},
		UserInfo{Name: "a1", Age: 7, Weight: 40},
		UserInfo{Name: "a1", Age: 8, Weight: 45},
	}
	_, err = cli.Collection.InsertMany(ctx, userInfos)
	if err != nil {
		fmt.Println(err)
	}

	// find one document
	one := UserInfo{}
	err = cli.Find(ctx, bson.M{"name": userInfo.Name}).One(&one)
	if err != nil {
		fmt.Println("user_info err")
	}
	fmt.Println("name", userInfo.Name)

	// find all 、sort and limit
	batch := []UserInfo{}
	cli.Find(ctx, bson.M{"age": 6}).Sort("weight").Limit(7).All(&batch)
	for _, v := range batch {
		fmt.Println("batch", v.Name)
	}
	
	// UpdateOne one
	err = cli.UpdateOne(ctx, bson.M{"name": "d4"}, bson.M{"$set": bson.M{"age": 7}})
	if err != nil {
		fmt.Println("UpdateOne err", err)
	}
	// UpdateAll
	result, err := cli.UpdateAll(ctx, bson.M{"age": 6}, bson.M{"$set": bson.M{"age": 10}})
	fmt.Println("UpdateAll", result)

	err = cli.Remove(ctx, bson.M{"age": 7})
	if err != nil {
		fmt.Println("Remove", err)
	}
}
