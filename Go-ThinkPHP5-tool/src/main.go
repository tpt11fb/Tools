package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const user_agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:98.0) Gecko/20100101 Firefox/98.0"
const content_type = "application/x-www-form-urlencoded"


func begin()  {
	fmt.Println(`
  _______ _     _       _    _____  _    _ _____    _______          _ 
 |__   __| |   (_)     | |  |  __ \| |  | |  __ \  |__   __|        | |
    | |  | |__  _ _ __ | | _| |__) | |__| | |__) |    | | ___   ___ | |
    | |  | '_ \| | '_ \| |/ /  ___/|  __  |  ___/     | |/ _ \ / _ \| |
    | |  | | | | | | | |   <| |    | |  | | |         | | (_) | (_) | |
    |_|  |_| |_|_|_| |_|_|\_\_|    |_|  |_|_|         |_|\___/ \___/|_|

																		--By Tptfb11`)
}

//接受参数
func agrs() (string,string,string) {
	var input_url = flag.String("url","","[-] plase input a url")
	var input_file = flag.String("f","","[-] -f input file")
	var input_module = flag.String("m","3","[-] choose module: 默认全部检测！\n'-m tp2' 检测thinkPHP2 rce\n'-m tp3' 检测thinkPHP3 log_rce\n'-m tp5' 检测thinkPHP5 sql+rce")
	flag.Parse()
	return *input_url, *input_file, *input_module
}

//处理url
func check_url(url string) string {
	if url[len(url)-1:] == "/"{
		url = url[0:len(url)-1]
		return url
	} else {
		return url
	}
}

func file_url(filepath string) []string {
	var urls[]string
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("[-] 请检查文件是否存在~.~")
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	reader := bufio.NewReader(file)
	for  {
		url, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		url = strings.TrimSpace(url)
		urls = append(urls,url)
	}
	return urls
}

func moudle(url string,m string)  {
	if (m == "tp5"){
		exec(url)
		debug_rce(url)
		get_rce(url)
		tp5_sql(url)
		others_rce(url)
		bybass(url)
	} else if (m == "tp3"){
		log_rce(url)
	} else if (m == "tp2") {
		tp2_rce(url)
	} else {
		tp2_rce(url)
		log_rce(url)
		exec(url)
		debug_rce(url)
		get_rce(url)
		tp5_sql(url)
		others_rce(url)
		bybass(url)
	}
}

func main() {
	log.SetOutput(io.Discard) //关闭打印Unsolicited response received on idle HTTP channel starting with 。。。服了这个
	begin()
	url, filepath, m := agrs()
	if (filepath != ""){
		urls := file_url(filepath)
		i := 0
		for _,url := range urls{
			i = i + 1
			url = check_url(url)
			fmt.Println("\n第"+ strconv.Itoa(i) +"个------------"+url)
			moudle(url,m)
		}
	}else if (url != "") {
		url = check_url(url)
		fmt.Println("\n------------"+url)
		moudle(url,m)
	}else {
		fmt.Println("[-] 请输入参数: -url 或者 -f  --^.^--")
	}
}