package main

import (
	"Audit/client"
	"encoding/json"
	"github.com/gin-gonic/gin"
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

const prefix = "audio::form::"

func Submit(c *gin.Context) {
	form := buildRequest(c)
	checkParams(form)
	if form.IsUpload == "true" {
		data, err := json.Marshal(form)
		if err != nil {
			log.Printf("marshal faild.err:%v", err)
			return
		}
		client.RedisClient.SetNX(prefix+form.Identifier, data, 0)
		c.String(http.StatusOK, "submit success")
		return
	}
	// todo:前段上传失败时上传MinIO

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
