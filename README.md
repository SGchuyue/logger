#golang实现的日志框架，提供日志保存、分割文件、压缩文件功能。
=====
#日志文件切割有两种类型：1为按日期切分。2为按日志大小切分。此功能实现按大小切分。
====
##当前版本状态：
------
指定网站URL访问返回信息，并保存到指定文件中；待完善存储到数据库中。<br>
输出日志格式默认如下：<br>
2020-09-10T19:56:30.288+0800    DEBUG   logger/mylog.go:46      正在尝试获得 https://github.com/ 的请求 <br>
2020-09-10T19:56:31.030+0800    INFO    logger/mylog.go:54      Success! statusCode = 200 OK for URL https://github.com/<br>
2020-09-10T19:56:31.030+0800    INFO    logger/mylog.go:55      成功获取！状态码=200 OK URL = https://github.com/<br>
2020-09-10T20:04:55.716+0800    DEBUG   logger/mylog.go:46      正在尝试获得 www.github.com 的请求<br>
2020-09-10T20:04:55.716+0800    ERROR   logger/mylog.go:50      Error fetching URL www.github.com : Error = Get www.github.com: unsupported protocol scheme ""<br>
2020-09-10T20:04:55.716+0800    ERROR   logger/mylog.go:51      获取URLwww.github.com时出错，错误信息为Get www.github.com: unsupported protocol scheme ""<br>
2020-09-10T20:04:55.716+0800    DEBUG   logger/mylog.go:46      正在尝试获得 http://github.com 的请求 <br>
2020-09-10T20:04:55.716+0800    ERROR   logger/mylog.go:50      Error fetching URL www.github.com : Error = Get www.github.com: unsupported protocol scheme ""<br>
2020-09-10T20:04:55.716+0800    ERROR   logger/mylog.go:51      获取URLwww.github.com时出错，错误信息为Get www.github.com: unsupported protocol scheme ""<br>
2020-09-10T20:21:00.062+0800    DEBUG   logger/mylog.go:47      正在尝试获得 https://github.com/ 的请求<br>
2020-09-10T20:21:00.684+0800    INFO    logger/mylog.go:55      Success! statusCode = 200 OK for URL https://github.com/<br>
2020-09-10T20:21:00.684+0800    INFO    logger/mylog.go:56      成功获取！状态码=200 OK URL = https://github.com/<br>

##使用方法：
------
通过修改`Filename`的位置改变日志文件存入位置。<br>
通过修改`MaxSize`的设定实现日志文件大小的设定。<br>
通过修改`MaxAges`的设定值实现文件的最大保存天数。<br>
通过修改`Compress`的类型选择是否压缩文件（默认为false）。<br>

##支持同一对象指定日志级别
----
参数说明：sLogger.Level 日志级别:<br>
如：sLogger.Infof("成功获取！状态码=%s URL = %s", resp.Status, url)<br>
 : sLogger.Errorf("获取URL%s时出错，错误信息为%s", url, err)<br>


