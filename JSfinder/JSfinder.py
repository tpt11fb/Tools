import re
from urllib.parse import urlparse
import requests
from bs4 import BeautifulSoup
import report

class JSfinder:
    def __init__(self,url="",cookie="",file_path="",html=""):
        self.cookie = cookie
        self.url = url
        self.file_path = file_path
        self.html = html

    def __del__(self):
        pass
    # 正则
    def extract_URL(self,JS):
        pattern_raw = r"""
          (?:"|')                               # Start newline delimiter
          (
            ((?:[a-zA-Z]{1,10}://|//)           # Match a scheme [a-Z]*1-10 or //
            [^"'/]{1,}\.                        # Match a domainname (any character + dot)
            [a-zA-Z]{2,}[^"']{0,})              # The domainextension and/or path
            |
            ((?:/|\.\./|\./)                    # Start with /,../,./
            [^"'><,;| *()(%%$^/\\\[\]]          # Next character can't be...
            [^"'><,;|()]{1,})                   # Rest of the characters can't be
            |
            ([a-zA-Z0-9_\-/]{1,}/               # Relative endpoint with /
            [a-zA-Z0-9_\-/]{1,}                 # Resource name
            \.(?:[a-zA-Z]{1,4}|action)          # Rest + extension (length 1-4 or action)
            (?:[\?|/][^"|']{0,}|))              # ? mark with parameters
            |
            ([a-zA-Z0-9_\-]{1,}                 # filename
            \.(?:php|asp|aspx|jsp|json|
                 action|html|js|txt|xml)             # . + extension
            (?:\?[^"|']{0,}|))                  # ? mark with parameters
          )
          (?:"|')                               # End newline delimiter
        """
        pattern = re.compile(pattern_raw, re.VERBOSE)
        result = re.finditer(pattern, str(JS))
        if result == None:
            return None
        js_url = []
        return [match.group().strip('"').strip("'") for match in result
                if match.group() not in js_url]
    # 请求
    def Extract_html(self,URL):
        header = {
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36",
            "Cookie": self.cookie}
        try:
            raw = requests.get(URL, headers=header, timeout=3, verify=False)
            raw = raw.content.decode("utf-8", "ignore")
            # print(raw)
            return raw
        except:
            return None
    #
    def process_url(self,URL, re_URL):
        black_url = ["javascript:"]  # Add some keyword for filter url.
        URL_raw = urlparse(URL)
        ab_URL = URL_raw.netloc
        host_URL = URL_raw.scheme
        if re_URL[0:2] == "//":
            result = host_URL + ":" + re_URL
        elif re_URL[0:4] == "http":
            result = re_URL
        elif re_URL[0:2] != "//" and re_URL not in black_url:
            if re_URL[0:1] == "/":
                result = host_URL + "://" + ab_URL + re_URL
            else:
                if re_URL[0:1] == ".":
                    if re_URL[0:2] == "..":
                        result = host_URL + "://" + ab_URL + re_URL[2:]
                    else:
                        result = host_URL + "://" + ab_URL + re_URL[1:]
                else:
                    result = host_URL + "://" + ab_URL + "/" + re_URL
        else:
            result = URL
        return result

    def find_last(self,string, point):
        positions = []
        last_position = -1
        while True:
            # print(type(string),type(point))
            position = string.find(point, last_position + 1)
            if position == -1: break
            last_position = position
            positions.append(position)
        return positions
    # 解析JS
    def find_by_url(self,url, html, js=False):
        if js == False:
            html = BeautifulSoup(html, "html.parser")
            html_scripts = html.findAll("script")
            script_array = {}
            script_temp = ""
            for html_script in html_scripts:
                script_src = html_script.get("src")
                if script_src == None:
                    script_temp += html_script.get_text() + "\n"
                else:
                    purl = self.process_url(url, script_src)
                    script_array[purl] = self.Extract_html(purl)
            script_array[url] = script_temp
            allurls = []
            for script in script_array:
                # print(script)
                temp_urls = self.extract_URL(script_array[script])
                if len(temp_urls) == 0: continue
                for temp_url in temp_urls:
                    allurls.append(self.process_url(script, temp_url))
            result = []
            for singerurl in allurls:
                url_raw = urlparse(url)
                domain = url_raw.netloc
                positions = self.find_last(domain, ".")
                # print(type(positions))
                miandomain = domain
                if len(positions) > 1: miandomain = domain[positions[-2] + 1:]
                # print(miandomain)
                suburl = urlparse(singerurl)
                subdomain = suburl.netloc
                # print(singerurl)
                if miandomain in subdomain or subdomain.strip() == "":
                    if singerurl.strip() not in result:
                        result.append(singerurl)
            return result
        return sorted(set(self.extract_URL(self.Extract_html(url)))) or None
    # 从JS中截取，域名
    def find_subdomain(self,urls, mainurl):
        url_raw = urlparse(mainurl)
        domain = url_raw.netloc # 获取到主域名
        positions = self.find_last(domain, ".")
        miandomain = domain
        if len(positions) > 1: miandomain = domain[positions[-2] + 1:]
        subdomains = []
        for url in urls:
            suburl = urlparse(url)
            subdomain = suburl.netloc
            # print(subdomain)
            if subdomain.strip() == "": continue
            if miandomain in subdomain:
                if subdomain not in subdomains:
                    subdomains.append(subdomain)
        return subdomains
    # 处理结果，日志输出
    def giveresult(self,urls, domian):
        if urls == None:
            return None
        print("Find " + str(len(urls)) + " URL:")
        content_url = ""
        content_subdomain = ""
        for url in urls:
            content_url += url + "\n"
            print(url)
        subdomains = self.find_subdomain(urls, domian)
        print("\nFind " + str(len(subdomains)) + " Subdomain:")
        for subdomain in subdomains:
            content_subdomain += subdomain + "\n"
            print(subdomain)
        # if args.outputurl != None:
        #     with open(args.outputurl, "a", encoding='utf-8') as fobject:
        #         fobject.write(content_url)
        #     print("\nOutput " + str(len(urls)) + " urls")
        #     print("Path:" + args.outputurl)
        # if args.outputsubdomain != None:
        #     with open(args.outputsubdomain, "a", encoding='utf-8') as fobject:
        #         fobject.write(content_subdomain)
        #     print("\nOutput " + str(len(subdomains)) + " subdomains")
        #     print("Path:" + args.outputsubdomain)
    # HTML报告
    def html_report(self,report_html):
        report.report(report_html)