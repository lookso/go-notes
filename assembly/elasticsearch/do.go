package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"os"
	"time"
)

type Es struct {
	Client *elastic.Client
	Ctx    context.Context
}

var MyIndexName = "my_es_first_index"

// 索引mapping定义，这里仿微博消息结构定义
const mapping = `{
 "mappings": {
   "properties": {
     "name": {
       "type": "keyword"
     },
     "country": {
       "type": "keyword"
     },
	 "age": {
       "type": "long"
     },
	 "date": {
       "type": "text"
     }
   }
 }
}`

type MappingIndex struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Age     int    `json:"age"`
	Date    string `json:"date"`
}

var doHost = "http://127.0.0.1:9200"

func main() {
	EsClient := NewEs()
	if err := EsClient.createIndex(); err != nil {
		fmt.Println("createIndex err", err)
	}
	if err := EsClient.insert(); err != nil {
		fmt.Println("insert err", err)
	}
}
func NewEs() *Es {
	var err error
	// 创建ES client用于后续操作ES
	client, err := elastic.NewClient(
		// 设置ES服务地址，支持多个地址
		elastic.SetURL(doHost),
		// 设置基于http base auth验证的账号和密码
		elastic.SetBasicAuth("admin", "123"),
		// 启用gzip压缩
		elastic.SetGzip(true),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置请求失败最大重试次数
		elastic.SetMaxRetries(5),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		// Handle error
		panic(err)
	}
	_, _, err = client.Ping(doHost).Do(context.Background())
	if err != nil {
		panic(err)
	}
	return &Es{
		Client: client,
		Ctx:    context.Background(),
	}
}

func (es *Es) createIndex() error {
	// 执行ES请求需要提供一个上下文对象
	es.Client.DeleteIndex(MyIndexName).Do(es.Ctx)
	// 首先检测下my_weibo_index索引是否存在
	exists, err := es.Client.IndexExists(MyIndexName).Do(es.Ctx)
	if err != nil {
		// Handle error
		return err
	}
	if !exists {
		// weibo索引不存在，则创建一个
		_, err := es.Client.CreateIndex(MyIndexName).BodyString(mapping).Do(es.Ctx)
		if err != nil {
			// Handle error
			return err
		}
	}
	return nil
}

func (es *Es) insert() error {
	// 创建创建一条数据
	msg := MappingIndex{Name: "jack", Country: "打酱油的一天", Age: 11, Date: "2021-12-12"}
	// 使用client创建一个新的文档
	put, err := es.Client.Index().
		Type(MyIndexName). // 设置索引名称
		Id("1"). // 设置文档id
		BodyJson(msg). // 指定前面声明的微博内容
		Do(es.Ctx) // 执行请求，需要传入一个上下文对象
	if err != nil {
		return err
	}
	fmt.Printf("文档Id %s, 索引名 %s\n", put.Id, put.Index)
	return nil
}
