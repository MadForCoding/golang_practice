package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

func main() {
	var wg sync.WaitGroup
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	partitionList, err := consumer.Partitions("abcd") //获得该topic所有的分区
	if err != nil {
		fmt.Println("Failed to get the list of partition:, ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("abcd", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		wg.Add(1)
		go func(sarama.PartitionConsumer) { //为每个分区开一个go协程去取值
			for msg := range pc.Messages() { //阻塞直到有值发送过来，然后再继续等待
				fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}
	wg.Wait()
	consumer.Close()

}
