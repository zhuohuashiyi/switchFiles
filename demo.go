package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
)

type Response struct {
	Anonymous string `json:"anonymous"`
	CheckCount int `json:"check_count"`
	FailCount int `json:"fail_count"`
	IsHTTPS bool `josn:"https"`
	LastStatus bool `json:"last_status"`
	LastTime string `json:"last_time"`
	Proxy string `json:"proxy"`
	Region string `json:"region"`
	Source string `json:"source"`
}

func main() {
	k := 0
	for k < 10 {
		resp, err := http.Get("http://demo.spiderpy.cn/get/")
		if err != nil {
			fmt.Println(err)
			continue
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var response Response
		err = json.Unmarshal(data, &response)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if !response.LastStatus || response.IsHTTPS{
			fmt.Println("the ip address is unavaiable, search for another")
			continue
		}
		uri := &url.URL{Host: response.Proxy}
		client := &http.Client{Transport: &http.Transport{
			Proxy: http.ProxyURL(uri),
		},}
    	req ,_ := http.NewRequest("GET", "http://baidu.com", nil)
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
		resp, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(resp.StatusCode)
		k++
	}
}