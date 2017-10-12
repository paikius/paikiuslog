package paikiuslog

import (
	"fmt"
	"log"
	"os"
	"time"
)

// SaveLog (fileLocation,text)
func SaveLog(fileLocation string, text string) {
	//get gopath location
	currentDate := time.Now().Local()
	path := os.Getenv("GOPATH") + fileLocation + "/logs/" + currentDate.Format("2006-01-02") + "-log.log"
	//check file exists
	_, err := os.Stat(path)

	if err != nil {
		log.Println("Log File Not Found")
		log.Println("Creating File .....")
		//create log file
		f, _ := os.Create(path)
		defer f.Close()
		f.WriteString(currentDate.Format("2006-01-02 15:04:05") + " => " + text + "\n\r")
	} else {
		fmt.Println("here")
		f, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0777)
		defer f.Close()
		f.WriteString(currentDate.Format("2006-01-02 15:04:05") + " => " + "demm" + "\n\r")
	}
}
