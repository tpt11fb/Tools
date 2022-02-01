# JSfinder

​	JSFinder是一款优秀的github开源工具，这款工具功能就是查找隐藏在js文件中的api接口和敏感目录，以及一些子域名。

	# 简单使用

## 简单爬取

```
python JSFinder.py -u http://www.mi.com

#这个命令会爬取 http://www.mi.com 这单个页面的所有的js链接，并在其中发现url和子域名
```

## 深度爬取

```
python JSFinder.py -u http://www.mi.com -d

#深入一层页面爬取JS，时间会消耗的更长,建议使用-ou 和 -os来指定保存URL和子域名的文件名
python JSFinder.py -u http://www.mi.com -d -ou mi_url.txt -os mi_subdomain.txt
```

## 批量指定URL/指定JS

指定URL：

```
python JSFinder.py -f text.txt
```

指定JS：

```
python JSFinder.py -f text.txt -j
```

可以用brupsuite爬取网站后提取出URL或者JS链接，保存到txt文件中，一行一个。
指定URL或JS就不需要加深度爬取，单个页面即可,等等，这可以去github上面看使用说明。

# 改造

​	为了方便简单使用和得出的数据结构更为直观，我对其进行如下的改造

![image-20220201133255573](/images/image-20220201133255573.png)

​	JSfinder改写成一个类，方便调用

![image-20220201133040787](/images/image-20220201133040787.png)

​	使用线程池加速爬取和解析

![image-20220201133418366](/images/image-20220201133418366.png)

​	导出报告，并对爬取到的js接口使用多线程进行验证，使拿到的数据跟精确（返回包是否200？是否404？）

![image-20220201133534652](/images/image-20220201133534652.png)

# 使用方式

## 简单爬取

```
python JSFinder.py -u http://www.mi.com

# 导出HTML
python JSFinder.py -u http://www.mi.com -html 1
```

## 批量爬取

```
python JSFinder.py -f text.txt

# 导出HTML
python JSFinder.py -f text.txt -html 1
```

## 使用截图

![image-20220201133859143](/images/image-20220201133859143.png)
