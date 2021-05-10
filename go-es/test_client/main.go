package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// 索引mapping定义，这里仿微博消息结构定义
const mapping = `
{
  "mappings": {
    "properties": {
      "user": {
        "type": "keyword"
      },
      "message": {
        "type": "text"
      },
      "image": {
        "type": "keyword"
      },
      "created": {
        "type": "date"
      },
      "tags": {
        "type": "keyword"
      },
      "location": {
        "type": "geo_point"
      },
      "suggest_field": {
        "type": "completion"
      }
    }
  }
}`

func main(){
	// 创建ES client用于后续操作ES
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9201"),
		// 设置基于http base auth验证的账号和密码
		elastic.SetBasicAuth("user", "secret"))
	if err != nil {
		// Handle error
		fmt.Printf("连接失败: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}

	ctx := context.Background()

	//检测索引是否存在
	exists, err := client.IndexExists("weibo").Do(ctx)
	if err != nil{
		panic(err)
	}
	if !exists{
		_, err := client.CreateIndex("weibo").BodyString(mapping).Do(ctx)
		if err != nil{
			panic(err)
		}
	}
}

