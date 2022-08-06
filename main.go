package main

import (
	"fmt"
	"log"
	"os"
	"path"
)


func main() {
	
	//her I´m extracting the actual folder path
	actualPath, _ := os.Getwd()

	//here we ar creating the path for the log file
	logFile := path.Join(actualPath, "logs.log")

	// The call to os.OpenFile() creates the log file for writing, 
	// if it does not already exist, or opens it for writing
	// by appending new data at the end of it (os.O_APPEND)
	//mor info about flags in https://pkg.go.dev/os#O_CREATE
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	//here we are verifying if there was errors when trying to open the file
	if err != nil {
		fmt.Println(err)
		return
	}

	//The defer keyword tells Go to execute the statement just before the current
	// function returns. This means that f.Close() is going to be executed just 
	//before main() returns.
	defer file.Close()

	logFromGo := log.New(file, "logFromGo: ", log.Ldate|log.Ltime|log.Llongfile)
	logFromGo.Println("primer log!")
    logFromGo.Println("otro log")

}
