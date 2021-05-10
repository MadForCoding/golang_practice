package es_CRUD

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"golang_practice/go-es/model"
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
	msg := model.Weibo{User: "olivere", Message: "打酱油的一天", Retweets: 0}
	ctx := context.Background()
	// 使用client创建一个新文档
	put, err := client.Index().Index("weibo").Id("1").BodyJson(msg).Do(ctx)

	if err != nil{
		panic(err)
	}
	fmt.Printf("文档Id %s, 索引名 %s\n", put.Id, put.Index)
}

func Get(){

}