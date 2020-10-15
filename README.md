 logger 是一个将日志分割、压缩、归档并写入指定文件的go包。
=====
### 此包提供了日志相关的五种常用功能：
+ **filename** 设置文件的存储位置
+ **maxsize** 设置切割文件的储存大小
+ **maxbackups** 保留旧日志文件的最大个数
+ **maxday** 保留旧日志文件的最大存留天数
+ **compress** 是否进行归档压缩
````
 func InitLogger(filename string, maxsize, maxbackups, maxday int, compress bool){}
````
-------
 example 使用方法：
------
1.导入需要日志包
> import "logger/logger"
  
2.进行初始化

````
func main() {
    // filename 设置存入日志文件路径及名称 字符串类型 （示例：./test.log）
    // maxsize 配置文件大小 int类型 （示例：2M）
    // maxbackups 配置旧文件保留个数 int类型 （示例： 3个）
    // maxday 配置日志文件保留最大天数 int类型 （示例： 5天）
    // compress 配置文件是否进行压缩 bool类型 （fasle不压缩，ture压缩） 
	logger.InitLogger("./test.log",2,3,5,true)
}
````
3.使用
> logger.Error("test error %s", "test")

----
测试包的功能
----
1.记录
````
// 测试logger日志包的功能
func TestInitLogger(t *testing.T) {
	logger.InitLogger("test.log",2,3,5,true)
	for i := 0; i < 50000; i++ {
		logger.Error("test  %s","test")
	}
	// PASS
	//ok      logger/logger   0.593s
}
````

---
时间可读性配置
---
*时间是以非人类可读的方式展示，例如1.572161051846623e+09*

*日常时间格式以年月日展示*
> encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
