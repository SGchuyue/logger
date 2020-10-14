package main

import "logger/logger"

//var sLogger *zap.SugaredLogger

func main() {
        // 初始化logger
        logger.InitLogger()
//      defer sLogger.Sync()
  //     logger.SimpleHttpGet("www.github.com")
   //    logger.SimpleHttpGet("http://github.com")
    //    logger.SimpleHttpGet("https://www")
//      zap.SugarLogger.Sync()
        log := logger.NewLogIn(text.log,2,3,6,false)
	logout := log.GetWriterLog()
}

