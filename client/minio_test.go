package client

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"testing"
)

func Test_MinIOInit(t *testing.T) {
	endpoint := "192.168.1.15:9000"

	// 初使化 minio client对象。
	opt := &minio.Options{}
	opt.Creds = credentials.New(MinioProvider{})
	minioClient, err := minio.New(endpoint, opt)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient初使化成功

	ctx := context.Background()
	// 创建一个叫mymusic的存储桶。
	bucketName := "audit"
	location := "2"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location, ObjectLocking: false})
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(ctx, bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	// 上传一个zip文件。
	objectName := "minio.go"
	filePath := "minio.go"
	contentType := "application/file"

	// 使用FPutObject上传一个zip文件。
	n, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %v\n", objectName, n)
}
