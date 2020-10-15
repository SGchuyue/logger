package main

import (
	"logger/logger"
)


func main() {
	logger.InitLogger("test.log",2,3,5,true)
	logger.Error("test error %s", "test")
	//logger(fmt.Println(time.Now()))
	//logger("hello")
	//logger("hello")
}

