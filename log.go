package main

import(
 "time"
 "fmt"
 "os"
 "log"
)

func CreateError(foo string){
	currentTime := time.Now()
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
	f, err := os.OpenFile("./errorlog", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    if _, err := f.Write([]byte(foo)); err != nil {
        log.Fatal(err)
	}
	
}
