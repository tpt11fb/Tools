# -*- coding=utf-8 -*-
#
import time, os
import requests
import threading
from urllib.parse import urlparse


class Template_mixin(object):
    HTML_TMPL = r"""
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <title>JSfinder测试报告</title>
            <link href="http://libs.baidu.com/bootstrap/3.0.3/css/bootstrap.min.css" rel="stylesheet">
            <h1 style="font-family: Microsoft YaHei">JSfinder测试报告</h1>
            <p class='attribute'><strong>测试结果 : </strong> %(value)s</p>
            <p class='attribute'><strong>筛选结果 : </strong><input type="text" id="filterName" placeholder="只限对状态的筛选" />
            <style type="text/css" media="screen">
        body  { font-family: Microsoft YaHei,Tahoma,arial,helvetica,sans-serif;padding: 20px;}
        </style>
        </head>
        <body>
                <script src="http://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>
                        <script type="text/javascript">
         $(function () {
          //设置列表查询
          $("#filterName").keyup(function () {
           // alert($(this).val())
           $("tr").stop().hide() //将tbody中的tr都隐藏
            .filter(":contains('"+($(this).val())+"')").show(); //，将符合条件的筛选出来
          });
         });
</script>
            %(table_ta)s
        </body>
        </html>"""

    TABLE_TABLE_TMPL = """
     <p class='attribute'><h2>Domain :  %(domain)s</h2><strong>JS数据共有%(js_num)s条</strong></p>
     <table id='result_table' class="table table-condensed table-bordered table-hover">
                <colgroup>
                    <col align='left' />
                    <col align='right' />
                    <col align='right' />
                    <col align='right' />
                </colgroup>
                <tr id='header_row' class="text-center success" style="font-weight: bold;font-size: 14px;">
                    <th>id</th>
                    <th>状态</th>
                    <th>URL</th>
                    <th>主域</th>
                    <th>备注</th>
                    <th>访问</th>
                </tr>
                %(table_tr)s
            </table>
    """
    TABLE_TMPL = """
        <tr class='failClass warning'>
            <td>%(id)s</td>
            <td>%(version)s</td>
            <td>%(step)s</td>
            <td>%(runresult)s</td>
            <td>%(runtime)s</td>
            <td>%(link)s</td>
        </tr>"""


def recheck(url):
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36",
    }
    try:
        res = requests.get(url, headers=headers)
        return res.status_code, res.text
    except:
        return "出错", "无HTML"


def get_domain(url):
    main = urlparse(url).netloc
    do = main.find(".") + 1
    domain = main[do:]
    return domain


def remarks(html):
    if html == "":
        return "这个页面是空的！"
    # elif key_word(html) != "":
    #     return "关键词{0}".format(key_word(html))
    else:
        return "无具体信息，请自行查看。"

def key_word(html):
    key_word = ["username", "passwd", "password", "admin", "login", "system", ""]
    k = ""
    for k in key_word:
        if k in html:
            k += k
    return k


def report(report_html):
    print("正在导出报告（数据越多，导出时间越长，保证了数据可靠，）----")
    table_table = ''
    html = Template_mixin()
    i = 0
    num = 0
    sum = 0
    for items in report_html:
        i += 1
        if (isinstance(items, list) & (len(items) >= 1)):
            table_tr0 = ""
            num += 1
            for item in items:
                if len(item) == 1:
                    break
                data = recheck(item)
                sum += 1
                table_td = html.TABLE_TMPL % dict(id=sum, version=data[0], step=item, runresult=urlparse(item).netloc,
                                                  runtime=remarks(data[1]),
                                                  link="<a href='%s' target='_blank''>点击访问</a>" % item)
                table_tr0 += table_td
            table_tab = html.TABLE_TABLE_TMPL % dict(domain=get_domain(items[0]), js_num=len(items), table_tr=table_tr0)
            table_table += table_tab
        else:
            pass
    total_str = '有 %s 个URL，共 %s 条数据，\n %s 条有效数据，' % (num, sum + i - num, sum)
    output = html.HTML_TMPL % dict(value=total_str, table_ta=table_table)
    # print('output',output)
    # 生成html报告
    filename = '{date}_TestReport.html'.format(date=time.strftime('%Y%m%d%H%M%S'))
    print("导出完毕：./report",filename)
    # 获取report的路径
    dir = os.path.join(os.getcwd(), 'report')
    filename = os.path.join(dir, filename)
    with open(filename, 'wb') as f:  # f.write(output)
        f.write(output.encode('utf8'))


def multi_thread_report(report_html,thread_num=5):
    threads = []
    for i in range(1,thread_num+1):
        threads.append(threading.Thread(target=report,args=(report_html,)))
    for thread in threads:
        thread.start()
    for thread in threads:
        thread.join()
    print('{date}_TestReport.html'.format(date=time.strftime('%Y%m%d%H%M%S')))

if __name__ == '__main__':
    report_html = ['https://www.baidu.com', ["https://www.baidu.com/nocache/fesplg/s.gif?log_type=sp",
                                             "https://www.baidu.com/nocache/fesplg/s.gif?log_type=hm&type=ssl&",
                                             "https://gt1.baidu.com/nocache/imgdata/sp613.gif?t="],
                   "https://www.tpt11fb.top", ["https://www.baidu.com/nocache/fesplg/s.gif?log_type=linksp",
                                               "https://sp1.baidu.com/5b1ZeDe5KgQFm2e88IuM_a/mwb2.gif?pid=",
                                               "https://www.baidu.com/wza/aria.js",
                                               "https://www.baidu.com/wza/aria.js?appid=c890648bf4dd00d05eb9751dd0548c30"]]
    multi_thread_report(report_html,5)