package client

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var MinioClient *minio.Client

const BucketName = "audit"

func InitMinIo() {
	endpoint := "127.0.0.1:9000"

	// 初使化 minio client对象。
	opt := &minio.Options{}
	opt.Creds = credentials.New(MinioProvider{})
	var err error
	MinioClient, err = minio.New(endpoint, opt)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", MinioClient) // minioClient初使化成功

}

type MinioProvider struct {
}

func (p MinioProvider) IsExpired() bool {
	return false
}

func (p MinioProvider) Retrieve() (credentials.Value, error) {
	return credentials.Value{AccessKeyID: "admin", SecretAccessKey: "password"}, nil
}
