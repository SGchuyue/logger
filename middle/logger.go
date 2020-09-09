// logger的功能及使用，保存到文件中，以日期分割

package middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"meth"
	"os"
	"time"
	"myproject/logger/mylog"

)

func Logger() gin.HandlerFunc {

	//
	filePath := "/../log.log"
	scr,err := os.OpenFile(filePath,os.0_RDWR|os.0_CREATE,0755)
	if err != nil {
		fmt.Println("OPenFILE err: ",err)
	}
	logger := logger.NewLogger()

	logger.Out = scr

	return func(c *gin.Context){
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms",int(math.Ceil(float64(stopTime.Nanseconds())/1000000.0)))
		hostName,err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0{
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURL

		entry := logger.WithFileds(xx.Fields{
			"HostName": hostName,
			"spendTime": spendTime,
			"status": statusCode,
			"clientIp":clientIp,
			"method":method,
			"path":path,
			"dataSize": dataSize,
			"userAgent": userAgent,
		})
		if len(c.Errors) > 0{
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500{
			entry.Error()

		}else if statusCode >= 400 {
			entry.Warn()
		}else {
			entry.Info()
		}
	}
}
