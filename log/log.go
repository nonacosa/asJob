package log

import (
	"log"
	"os"
)


func Log() {
	//os.Create("job.log")
	logFile,err := os.OpenFile("job.log",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error!")
	}
	debugLog := log.New(logFile,"--[Debug]--",log.Ltime)
	debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	debugLog.Println("A debug message here")

}