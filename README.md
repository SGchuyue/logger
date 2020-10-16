 logger 是一个基于zap、zapcore及Lumberjack等优秀开源项目将日志记录、分割、压缩并写入指定文件等功能封装起来的go日志记录器。
=====
### 此包提供了日志相关的五种常用功能：
+ 自定义文件存储位置，将其打印到文件中而非终端工作区。
+ 日志切割，根据工作需求自由设置文件的存储大小。
+ 日志存留，以文件个数和日期更迭选择日志文件的存留时间。
+ 日志压缩，对保留的日志文件压缩节省空间。
+ 可自由选择不同级别的日志打印。

 example 使用方法：
------
下载包的依赖
>go get -u "github.com/SGchuyue/logger/logger"

导入日志包进行初始化后即可使用
````
package main

import (
	"github.com/SGchuyue/logger/logger"
)

func main() {
    // filename 设置存入日志文件路径及名称 字符串类型 （示例：./test.log）
    // maxsize 配置文件大小 int类型 （示例：2M）
    // maxbackups 配置旧文件保留个数 int类型 （示例： 3个）
    // maxday 配置日志文件保留最大天数 int类型 （示例： 5天）
    // compress 配置文件是否进行压缩 bool类型 （fasle不压缩，ture压缩） 
    // 对日志包的初始化
    logger.InitLogger("./test.log",2,3,5,true)
    // 对日志包的使用
    logger.Error("error test")
    logger.Errorf("test errorf %s","test")
    logger.Debug("debug test")
    logger.Debugf("test debugf %s","test")
    logger.Info("info test")
    logger.Infof("test infof %s","test")
    logger.Warn("warn test")
    logger.Warnf("test warnf %s","test")
}
````

测试功能实现
---

````
// 测试logger包的功能
func TestInitLogger(t *testing.T) {
	logger.InitLogger("test.log",2,3,5,true)
	for i := 0; i < 50000; i++ {
    	    logger.Errorf("test errorf %s","test")
    	    logger.Debug("debug test")
    	    logger.Debugf("test debugf %s","test")
    	    logger.Info("info test")
    	    logger.Infof("test infof %s","test")
    	    logger.Warn("warn test")
    	    logger.Warnf("test warnf %s","test")
	}
}
````