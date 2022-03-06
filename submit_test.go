package main

import (
	"Audit/client"
	"encoding/json"
	"fmt"
	"github.com/alicebob/miniredis/v2" //nolint:typecheck
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSubmit(t *testing.T) {
	tests := []AudioForm{
		{"echo", "110", "522", "test", "test", "true", nil},
	}
	mini := miniredis.RunT(t) //nolint:typecheck
	defer mini.Close()
	client.RedisClient = redis.NewClient(&redis.Options{Addr: mini.Addr()})
	for _, test := range tests {
		gin.SetMode(gin.TestMode)
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		req, _ := http.NewRequest("POST", "/submit", nil)
		req.PostForm = url.Values{}
		c.Request = req
		js, _ := json.Marshal(test)
		mp := make(map[string]string)
		err := json.Unmarshal(js, &mp)
		if err != nil {
			t.Errorf("unmarshal error: %v", err)
		}
		for key, value := range mp {
			c.Request.PostForm.Add(key, value)
		}
		Submit(c)
		if res, _ := mini.Get(prefix + test.Identifier); res == "" {
			t.Errorf("submit can't work.")
		} else {
			fmt.Println(res)
		}
	}
}

func Test_Demo(t *testing.T){
	s := "赞"
	fmt.Println(len(s))
}
