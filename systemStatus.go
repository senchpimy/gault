package main

import (
	"os/exec"
	"strings"
)

type Status struct{
	Service string
	Ennabled bool
	Status string
	Uptime string

}

type SystemStatusStruct struct{
Services []Status
}

func StatusFtp()(foo string){
    cmd2 := exec.Command("systemctl", "status", "vsftpd")
    out, _ := cmd2.CombinedOutput()
    out2:=string(out)
    // if err2 != nil {log.Fatal(err2)}
    return out2
 }

func StatusNfs()(foo string){
    cmd2 := exec.Command("systemctl", "status", "nfs-server")
    out, _ := cmd2.CombinedOutput()
    out2:=string(out)
    // if err2 != nil {log.Fatal(err2)}
    return out2
 }

func StatusSSH()(foo string){
    cmd2 := exec.Command("systemctl", "status", "sshd")
    out, _ := cmd2.CombinedOutput()
    out2:=string(out)
    // if err2 != nil {log.Fatal(err2)}
    return out2
 }

func StatusSmb()(foo string){
    cmd2 := exec.Command("systemctl", "status", "smb")
    out, _ := cmd2.CombinedOutput()
    out2:=string(out)
    // if err2 != nil {log.Fatal(err2)}
    return out2
 }

 func StatusFormater(StatusInputPrev string)(foo Status){
	StatusInput:=strings.Split(StatusInputPrev,"\n")[0:3]
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

func SystemStatus()(foo SystemStatusStruct){
	services:=make([]Status,4)
	services[0]=StatusFormater(StatusNfs())
	services[1]=StatusFormater(StatusSmb())
	services[2]=StatusFormater(StatusFtp())
	services[3]=StatusFormater(StatusSSH())
	var bar SystemStatusStruct
	bar.Services=services
	return bar
}
