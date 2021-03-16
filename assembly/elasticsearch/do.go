package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

type Es struct {
	Client *elastic.Client
	Ctx    context.Context
}

var myIndexName = "my_es_first_index"

// 索引mapping定义，这里仿微博消息结构定义
const mapping = `{
 "mappings": {
   "properties": {
     "name": {
       "type": "text"
     },
     "country": {
       "type": "text"
     },
	 "age": {
       "type": "long"
     }
   }
 }
}`

type MappingIndex struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Age     int64  `json:"age"`
}

var doHost = "http://127.0.0.1:9200"

func main() {
	EsClient := NewEs()
	if err := EsClient.CreateIndex(); err != nil {
		fmt.Println("createIndex err", err)
	}
	if err := EsClient.Insert(); err != nil {
		fmt.Println("insert err", err)
	}
	if err := EsClient.Search(); err != nil {
		fmt.Println("search err", err)
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
	esVersion, err := client.ElasticsearchVersion(doHost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esVersion) // Elasticsearch version 7.10.0
	return &Es{
		Client: client,
		Ctx:    context.Background(),
	}
}

func (es *Es) CreateIndex() error {
	_, err := es.Client.DeleteIndex(myIndexName).Do(es.Ctx)
	// 执行ES请求需要提供一个上下文对象
	// 首先检测下my_weibo_index索引是否存在
	exists, err := es.Client.IndexExists(myIndexName).Do(es.Ctx)
	fmt.Println("index exists", exists)
	if err != nil {
		// Handle error
		return err
	}
	if !exists {
		// weibo索引不存在，则创建一个
		_, err := es.Client.CreateIndex(myIndexName).BodyString(mapping).Do(es.Ctx)
		if err != nil {
			// Handle error
			return err
		}
	}
	return nil
}

func (es *Es) Insert() error {
	// 使用client创建一个新的文档
	for docId := 1; docId <= 100; {
		// 创建创建一条数据
		docIdStr := strconv.Itoa(docId)
		msg := MappingIndex{Name: "jack" + docIdStr, Country: "打酱油的一天", Age: int64(10 + docId)}
		put, err := es.Client.Index().
			Index(myIndexName). // 设置索引名称
			Type("_doc"). // 默认使用_doc type
			Id(docIdStr). // 设置文档id
			BodyJson(msg). // 指定前面声明的微博内容
			Do(es.Ctx) // 执行请求，需要传入一个上下文对象
		if err != nil {
			return err
		}
		fmt.Printf("文档Id %s, 索引名 %s\n", put.Id, put.Index)
		docId = docId + 1
	}

	return nil
}

func (es *Es) Search() error {
	get, err := es.Client.Get().
		Index(myIndexName). // 指定索引名
		Type("_doc").
		Id("1"). // 设置文档id
		Do(es.Ctx) // 执行请求
	if err != nil {
		return err
	}
	if get.Found {
		fmt.Printf("search 文档id=%s 版本号=%d 索引名=%s\n", get.Id, get.Version, get.Index)
	}
	// 创建term查询条件，用于精确查询
	termQuery := elastic.NewMatchQuery("name", "jack11")
	searchResult, err := es.Client.Search().
		Index(myIndexName). // 设置索引名
		Query(termQuery). // 设置查询条件
		Sort("age", true). // 设置排序字段，根据age字段升序排序，第二个参数false表示逆序
		From(0). // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(10). // 设置分页参数 - 每页大小
		Pretty(true). // 查询结果返回可读性较好的JSON格式
		Do(es.Ctx) // 执行请求
	if err != nil {
		return err
	}
	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var b1 MappingIndex
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(b1)) {
			// 转换成MappingIndex对象
			if t, ok := item.(MappingIndex); ok {
				fmt.Println(t.Name, t.Country, t.Age)
			}
		}
	}
	return nil
}

// 参考文档:
// https://www.tizi365.com/archives/850.html
