//package main
//
//import (
//	"context"
//	"fmt"
//	"github.com/olivere/elastic"
//)
//
//var weibo = `{
//    "name": "wali",
//    "country": "Chian",
//    "age": 30,
//    "date": "1987-03-07"
//    }`
//
//type Es struct {
//	Client *elastic.Client
//	Ctx    context.Context
//}
//
//var WeiboIndexName = "my_weibo_index"
//
//func main() {
//	EsClient := NewEs()
//	if err := EsClient.createIndex(); err != nil {
//		fmt.Println("createIndex err", err)
//	}
//	if err := EsClient.insert(); err != nil {
//		fmt.Println("insert err", err)
//	}
//}
//func NewEs() *Es {
//	var err error
//	// 创建ES client用于后续操作ES
//	client, err := elastic.NewClient(
//		// 设置ES服务地址，支持多个地址
//		elastic.SetURL("http://127.0.0.1:9200"),
//		// 设置基于http base auth验证的账号和密码
//		elastic.SetBasicAuth("admin", "123"))
//	if err != nil {
//		// Handle error
//		panic(err)
//	}
//	return &Es{
//		Client: client,
//		Ctx:    context.Background(),
//	}
//}
//
//func (es *Es) createIndex() error {
//
//	// 执行ES请求需要提供一个上下文对象
//	es.Client.DeleteIndex(WeiboIndexName).Do(es.Ctx)
//	// 首先检测下my_weibo_index索引是否存在
//	exists, err := es.Client.IndexExists(WeiboIndexName).Do(es.Ctx)
//	if err != nil {
//		// Handle error
//		return err
//	}
//	if !exists {
//		// weibo索引不存在，则创建一个
//		_, err := es.Client.CreateIndex(WeiboIndexName).Do(es.Ctx)
//		if err != nil {
//			// Handle error
//			return err
//		}
//	}
//	return nil
//}
//func (es *Es) insert() error {
//	// 创建创建一条微博
//	//tags := []string{"jack"}
//	//msg1 := Weibo{User: "jack", Message: "打酱油的一天", Image: "this image url", Created: time.Now(), Tags: tags, Location: "", Suggest: nil}
//
//	// 使用client创建一个新的文档
//	put1, err := es.Client.Index().
//		Index(WeiboIndexName). // 设置索引名称
//		Id("1"). // 设置文档id
//		BodyJson(weibo). // 指定前面声明的微博内容
//		Do(es.Ctx) // 执行请求，需要传入一个上下文对象
//	if err != nil {
//		// Handle error
//		return err
//	}
//	fmt.Printf("文档Id %s, 索引名 %s\n", put1.Id, put1.Index)
//	return nil
//}

package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

//创建索引：
func main() {
	Client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	fmt.Println(Client, err)
	name := "people2"
	Client.CreateIndex(name).Do(context.Background())

	data := `{
   "name": "wali",
   "country": "Chian",
   "age": 30,
   "date": "1987-03-07"
   }`
	_, err = Client.Index().Index(name).Type("man1").Id("1").BodyJson(data).Do(context.Background())

}

//
////插入数据
//func main(){

//}
//
//查找数据： //通过id查找
//func main(){
//	Client, err := elastic.NewClient(elastic.SetURL("http://192.168.7.6:9200"))
//	fmt.Println(Client, err)
//	name := "people2"
//	get, err := Client.Get().Index(name).Type("man1").Id("1").Do(context.Background())
//	fmt.Println(get, err)
//
//}
//
//
//
////修改
//func main() {
//	Client, err := elastic.NewClient(elastic.SetURL("http://192.168.7.6:9200"))
//	res, err := client.Update().
//		Index("megacorp").
//		Type("employee").
//		Id("2").
//		Doc(map[string]interface{}{"age": 88}).
//		Do(context.Background())
//	if err != nil {
//		println(err.Error())
//	}
//	fmt.Printf("update age %s\n", res.Result)
//
//}
