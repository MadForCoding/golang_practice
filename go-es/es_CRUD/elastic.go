package es_CRUD

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"golang_practice/go-es/model"
	"reflect"
)

var client = NewClient()

func NewClient() *elastic.Client{
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9201"),
		elastic.SetBasicAuth("user", "secret"))
	if err != nil{
		fmt.Printf("连接失败: %v\n", err)
		panic(err)
		return nil
	}
	return client
}


func Insert(){
	msg := model.Weibo{User: "olivere", Message: "blabla,pending,nage", Retweets: 0}
	ctx := context.Background()
	// 使用client创建一个新文档
	put, err := client.Index().Index("weibo").Id("3").BodyJson(msg).Do(ctx)

	if err != nil{
		panic(err)
	}
	fmt.Printf("文档Id %s, 索引名 %s\n", put.Id, put.Index)
}

func Get(){
	// 根据id查询文档
	get1, err := client.Get().
		Index("weibo"). // 指定索引名
		Id("1"). // 设置文档id
		Do(context.Background()) // 执行请求
	if err != nil {
		// Handle error
		panic(err)
	}
	if get1.Found {
		fmt.Printf("文档id=%s 版本号=%d 索引名=%s\n", get1.Id, get1.Version, get1.Index)
	}

	// 手动将文档内容转换成go struct对象
	msg2 := model.Weibo{}
	// 提取文档内容，原始类型是json数据
	data, _ := get1.Source.MarshalJSON()
	// 将json转成struct结果
	json.Unmarshal(data, &msg2)
	// 打印结果
	fmt.Println(msg2)
}

func GetIndexInfo() {
	get1, err := client.IndexGet("weibo").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(get1["weibo"].Mappings)
}

func Query() {
	q := elastic.NewBoolQuery()
	q.Must(elastic.NewMatchQuery("message", "zhegeeee   , "))

	res, err := client.Search("weibo").Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}
	var rList []*model.Weibo
	for _, item := range res.Each(reflect.TypeOf((*model.Weibo)(nil))) {
		rList = append(rList, item.(*model.Weibo))
		fmt.Println(item.(*model.Weibo))
	}
}