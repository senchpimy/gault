package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
    	"os/exec"
)

func WriteToFile(Texto string, File string) {
	// Read Write Mode
	file, err := os.OpenFile(File, os.O_RDWR, 0644)
	
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	
	len, err := file.WriteString(Texto) // Write at 0 beginning
	if err != nil {
		log.Fatalf("failed writing to file: %s %s", err,len)
	}
}

func ReadFile(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Content: %s", data)	
}

func ExistSambaConf()  {
	if _, err := os.Stat("/etc/samba/smb.conf"); errors.Is(err, os.ErrNotExist) {
		log.Fatal("Samba Config File 'smb.conf' does not exist")
		}
}
