package main

import (
	"Audit/client"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"log"
	"mime/multipart"
	"net/http"
)

type AudioForm struct {
	Name       string `json:"name"`
	Telephone  string `json:"telephone"`
	Identifier string `json:"identifier"`
	BucketName string `json:"bucket_name"`
	ObjectName string `json:"object_name"`
	IsUpload   string `json:"is_upload"`
	UploadFile *multipart.FileHeader
}

const prefix = "audit::form::"

func Submit(c *gin.Context) {
	form := buildRequest(c)
	checkParams(form)
	if form.IsUpload == "true" {
		json, _ := json.Marshal(form)
		client.RedisClient.SetNX(prefix+form.Identifier, json, 0)
		c.String(http.StatusOK, "submit success")
		return
	}
	// 前段上传失败时上传MinIO
	file, err := form.UploadFile.Open()
	defer file.Close()
	if err != nil {
		log.Printf("OPen File failed.err:%v\n", err)
		c.String(http.StatusOK, "Internal Error.")
		return
	}
	opts := minio.PutObjectOptions{ContentType: form.UploadFile.Header.Get("Content-Type")}
	info, err := client.MinioClient.PutObject(c, client.BucketName, form.Identifier+"_"+form.UploadFile.Filename, file, form.UploadFile.Size, opts)
	if err != nil {
		log.Printf("Upload file failed.err:%v", err)
		c.String(http.StatusOK, "Internal Error.")
		return
	}
	form.BucketName = info.Bucket
	form.ObjectName = info.Key
	json, _ := json.Marshal(form)
	c.String(http.StatusOK, "submit success")
	client.RedisClient.SetNX(prefix+form.Identifier, json, 0)
}

func checkParams(form AudioForm) {
	if form.Name == "" {
		log.Printf("name is empty.")
		return
	}
	if form.Telephone == "" {
		log.Printf("telephone is empty.")
		return
	}
	if form.Identifier == "" {
		log.Printf("invalid identifier.")
		return
	}
}

func buildRequest(c *gin.Context) AudioForm {
	form := AudioForm{}
	var err error
	form.Name = c.DefaultPostForm("name", "")
	form.Telephone = c.DefaultPostForm("telephone", "")
	form.Identifier = c.DefaultPostForm("identifier", "")
	form.BucketName = c.DefaultPostForm("bucket_name", "")
	form.ObjectName = c.DefaultPostForm("object_name", "")
	form.IsUpload = c.DefaultPostForm("is_upload", "false")
	if form.IsUpload == "false" {
		form.UploadFile, err = c.FormFile("upload_file")
		if err != nil {
			log.Printf("[buildRequest]file upload failed.err:%v", err)
		}
	}
	return form
}
