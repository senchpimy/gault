package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func CreateError(foo string){
	currentTime := time.Now()
	date:=currentTime.Format("2006-01-02 15:04:05")
	f, err := os.OpenFile("./errorlog", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    if _, err := f.Write([]byte(date+": "+foo+"\n")); err != nil {
        log.Fatal(err)
	}
	fmt.Println(date+": "+foo)
}
