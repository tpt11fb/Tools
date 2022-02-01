#!/usr/bin/env python"
# coding: utf-8
# By Tptfb11
# http://tpt11fb.top/Tptfb11

import argparse
import concurrent.futures
import sys
import urllib3
import JSfinder

def parse_args():
    parser = argparse.ArgumentParser(epilog='\tExample: \r\npython ' + sys.argv[0] + " -u http://www.baidu.com")
    parser.add_argument("-u", "--url", help="The website")
    parser.add_argument("-c", "--cookie", help="The website cookie")
    parser.add_argument("-f", "--file", help="The file contains url or js")
    parser.add_argument("-ou", "--outputurl", help="Output file name. ")
    parser.add_argument("-os", "--outputsubdomain", help="Output file name. ")
    parser.add_argument("-j", "--js", help="Find in js file", action="store_true")
    parser.add_argument("-d", "--deep",help="Deep find", action="store_true")
    parser.add_argument("-html", "--html_thread",type=int,help="The number of threads for exporting HTML reports. The default is 5")
    return parser.parse_args()

# 请求线程池
def thread_pool_askUrl(relative,urls):
    with concurrent.futures.ThreadPoolExecutor() as pool:
        htmls = pool.map(relative.Extract_html, urls)
        htmls = list(zip(urls, htmls))
        # 取出元素
        # for url, html in htmls:
        #     print(url, len(html))
        # print(htmls)
        return htmls
# 解析线程池
def thread_pool_deelData(relative,htmls):
    with concurrent.futures.ThreadPoolExecutor() as pool:
        futures = {}
        datas=[]
        for url, html in htmls:
            future = pool.submit(relative.find_by_url, url,html)
            futures[future] = url
        # 输出结果
        for future, url in futures.items():
            try:
                datas.append(url)
                datas.append(future.result())
                jsfinder.giveresult(future.result(),url)
            except:
                print("Fail to access："+url)
        return datas

def  file_open(file_path):
    with open(file_path, "r") as fobject:
        links = fobject.read().split("\n")
    if links == []: return None
    print("ALL Find " + str(len(links)) + " links")
    return links

if __name__ == '__main__':
    urllib3.disable_warnings()
    args = parse_args()
    urls = []
    if args.file != None:
        urls = file_open(args.file)
    elif args.url != None:
        urls.append(args.url)
    else:
        print("[-]error：At least one URL is required！\n[-]Multiple URLs, please use the parameter -f")
        exit()
    print(urls)
    jsfinder = JSfinder.JSfinder(cookie=args.cookie,html=args.html_thread)
    htmls = thread_pool_askUrl(jsfinder,urls)
    report_html = thread_pool_deelData(jsfinder,htmls)
    if args.html_thread != None:
        jsfinder.html_report(report_html)
    else:
        print("[+]-html can also generate HTML reports！")

