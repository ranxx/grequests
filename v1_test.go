package grequests_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ranxx/grequests"
)

func TestGet(t *testing.T) {

	resp := make([]byte, 0, 2000)

	fmt.Println(grequests.Get(context.TODO(), "http://www.baidu.com", nil, &resp, grequests.SetDecoder(func(data []byte, out interface{}) error {
		resp = data
		return nil
	})))

	fmt.Println(string(resp))
}

func TestPost(t *testing.T) {

	req := struct {
		Laddr []struct {
			Ip   string `json:"ip"`
			Port int    `json:"port"`
		} `json:"laddr"`
		Raddr struct {
			Ip   string `json:"ip"`
			Port int    `json:"port"`
		} `json:"raddr"`
	}{
		Laddr: []struct {
			Ip   string "json:\"ip\""
			Port int    "json:\"port\""
		}{
			{Port: 3337},
		},
		Raddr: struct {
			Ip   string "json:\"ip\""
			Port int    "json:\"port\""
		}{
			Port: 3333,
		},
	}

	resp := make([]byte, 0, 2000)

	fmt.Println(grequests.Post(context.TODO(), "http://localhost:12351/transfer/tcp", req, &resp, grequests.SetDecoder(func(data []byte, out interface{}) error {
		resp = data
		return nil
	})))

	fmt.Println("resp", string(resp))
}

func TestPostV2(t *testing.T) {

	resp := []struct {
		ID      int64 `json:"id"`
		Network int   `json:"network"`
		Laddr   struct {
			Ip   string `json:"ip"`
			Port int    `json:"port"`
		} `json:"laddr"`
		Raddr struct {
			Ip   string `json:"ip"`
			Port int    `json:"port"`
		} `json:"raddr"`
	}{}

	fmt.Println(grequests.Get(context.TODO(), "http://localhost:12351/transfer", nil, &resp))

	data, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println(string(data))
}
