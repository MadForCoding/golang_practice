package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	//uploadFile()
	//listFile()
	downloadFile()
}

func newClient() (*s3.S3, *session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(s3Cfg.AccessKey, s3Cfg.SecretKey, ""),
		Region:           aws.String("sg"),
		Endpoint:         aws.String(s3Cfg.Host),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
		LogLevel:         aws.LogLevel(aws.LogDebug),
	})
	if err != nil {
		return nil, nil, err
	}

	svc := s3.New(sess)
	svc.AddDebugHandlers()

	return svc, sess, nil
}

var s3Cfg = struct {
	Host      string
	Bucket    string
	AccessKey string
	SecretKey string
}{
	Host:      "xx",
	Bucket:    "xx",
	AccessKey: "xx",
	SecretKey: "xx",
}
