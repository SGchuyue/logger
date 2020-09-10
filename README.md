# logger
=====

#封装一个函数复用，实现日志文件的记录、存储、切割、压缩功能。
====
效果如下：
-------
2020-09-10T19:56:30.288+0800    DEBUG   logger/mylog.go:46      正在尝试获得 https://github.com/ 的请求 <br>
2020-09-10T19:56:31.030+0800    INFO    logger/mylog.go:54      Success! statusCode = 200 OK for URL https://github.com/<br>
2020-09-10T19:56:31.030+0800    INFO    logger/mylog.go:55      成功获取！状态码=200 OK URL = https://github.com/<br>
2020-09-10T20:04:55.716+0800    DEBUG   logger/mylog.go:46      正在尝试获得 www.github.com 的请求<br>
2020-09-10T20:04:55.716+0800    ERROR   logger/mylog.go:50      Error fetching URL www.github.com : Error = Get www.github.com: unsupported protocol scheme ""<br>
2020-09-10T20:04:55.716+0800    ERROR   logger/mylog.go:51      获取URLwww.github.com时出错，错误信息为Get www.github.com: unsupported protocol scheme ""<br>
2020-09-10T20:04:55.716+0800    DEBUG   logger/mylog.go:46      正在尝试获得 http://github.com 的请求
2020-09-10T20:04:55.975+0800    ERROR   logger/mylog.go:50      Error fetching URL http://github.com : Error = Get http://github.com: read tcp 192.168.1.6:36254->13.250.177.223:80: read: connection reset by peer
2020-09-10T20:04:55.975+0800    ERROR   logger/mylog.go:51      获取URLhttp://github.com时出错，错误信息为Get http://github.com: read tcp 192.168.1.6:36254->13.250.177.223:80: read: connection reset by peer
2020-09-10T20:20:59.913+0800    DEBUG   logger/mylog.go:47      正在尝试获得 www.github.com 的请求
2020-09-10T20:20:59.913+0800    ERROR   logger/mylog.go:51      Error fetching URL www.github.com : Error = Get www.github.com: unsupported protocol scheme ""
2020-09-10T20:20:59.913+0800    ERROR   logger/mylog.go:52      获取URLwww.github.com时出错，错误信息为Get www.github.com: unsupported protocol scheme ""
2020-09-10T20:20:59.913+0800    DEBUG   logger/mylog.go:47      正在尝试获得 http://github.com 的请求
2020-09-10T20:21:00.062+0800    ERROR   logger/mylog.go:51      Error fetching URL http://github.com : Error = Get http://github.com: read tcp 192.168.1.6:36412->13.250.177.223:80: read: connection reset by peer
2020-09-10T20:21:00.062+0800    ERROR   logger/mylog.go:52      获取URLhttp://github.com时出错，错误信息为Get http://github.com: read tcp 192.168.1.6:36412->13.250.177.223:80: read: connection reset by peer
2020-09-10T20:21:00.062+0800    DEBUG   logger/mylog.go:47      正在尝试获得 https://github.com/ 的请求
2020-09-10T20:21:00.684+0800    INFO    logger/mylog.go:55      Success! statusCode = 200 OK for URL https://github.com/
2020-09-10T20:21:00.684+0800    INFO    logger/mylog.go:56      成功获取！状态码=200 OK URL = https://github.com/
 
使用方法：
------
