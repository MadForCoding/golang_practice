package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func listFile() {
	svc, _, err := newClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	params := &s3.ListObjectsInput{
		Bucket: aws.String(s3Cfg.Bucket),
	}

	resp, err := svc.ListObjects(params)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func uploadFile() {
	_, sess, err := newClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := "data_part.csv"
	file := "./amazon_s3/" + fileName

	fp, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	param := &s3manager.UploadInput{
		Bucket:      aws.String(s3Cfg.Bucket),
		Key:         aws.String(fileName),
		Body:        fp,
		ContentType: aws.String("application/csv"),
	}

	uploader := s3manager.NewUploader(sess)
	resp, err := uploader.Upload(param)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func downloadFile() {
	_, sess, err := newClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpFile := "./amazon_s3/tmp.csv"
	_ = os.Remove(tmpFile)
	fp, err := os.Create(tmpFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(fp,
		&s3.GetObjectInput{
			Bucket: aws.String(s3Cfg.Bucket),
			Key:    aws.String("data_part.csv"),
		})
	fmt.Printf("numBytes: %d, err: %+v\n", numBytes, err)
}
