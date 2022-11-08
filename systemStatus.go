package main

import (
//	"encoding/json"
	"fmt"
//	"io/ioutil"
	// "log"
//	"math/rand"
//	"os"
	"os/exec"
//	"strconv"
	"strings"
//	"time"
)

type Status struct{
	Service string
	Ennabled bool
	Status string
	Uptime string

}

func StatusFtp()(foo []string){
    cmd2 := exec.Command("systemctl", "status", "vsftpd")
    out, _ := cmd2.CombinedOutput()
    out2:=string(out)
    // if err2 != nil {log.Fatal(err2)}
    return strings.Split(out2,"\n")[0:3]
 }

func StatusNfs()(foo []string){
    cmd2 := exec.Command("systemctl", "status", "nfs-server")
    out, _ := cmd2.CombinedOutput()
    out2:=string(out)
    // if err2 != nil {log.Fatal(err2)}
    return strings.Split(out2,"\n")[0:3]
 }

func StatusSmb()(foo []string){
    cmd2 := exec.Command("systemctl", "status", "smb")
    out, _ := cmd2.CombinedOutput()
    out2:=string(out)
    // if err2 != nil {log.Fatal(err2)}
    return strings.Split(out2,"\n")[0:3]
 }

 func StatusFormater(StatusInput []string)(foo Status){
	foo.Service=strings.Split(StatusInput[0]," ")[1]

	if len(strings.Split(StatusInput[2]," "))>10{
		SecondLine:=strings.Split(StatusInput[2]," ")
		foo.Uptime=SecondLine[10]+" "+SecondLine[11]
		foo.Status=SecondLine[6]
	}else{
		foo.Uptime="Disabled"
		SecondLine:=strings.Split(StatusInput[2]," ")
		foo.Status=SecondLine[6]
	}

	IsEnabled:=strings.Split(StatusInput[1],";")[1]
	if IsEnabled==" enabled"{
		foo.Ennabled=true
	}else{
		foo.Ennabled=false
	}
	return foo
}

func main()  {
StatusFormater(StatusNfs())
}
