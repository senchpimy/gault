package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"encoding/json"
)

//sudo blkid:UUID
//sudo fdisk -l: mucha info
//lsblk: Discos
//df: sistemas ya montados

type Disk_DF struct{
	Filesystem string
	Mem int
	Used int
	Avaible int
	UsePercent string
	Mount string
}

type Format_DF struct{
Title string
Todos []Disk_DF
}

type Format_lsblk struct{
Title string
Todos System_lsblk
}

type Disk_lbslk struct{
	Name string 
	Size string 
	Type string
}

type System_lsblk struct{
	Blockdevices []Disk_lbslk
}

func GetInfoSystem() (bar [][]string) {
    out, err := exec.Command("df").Output()
    if err != nil {
        log.Fatal(err)
    }
    dfOutput:=strings.Fields(string(out))[7:]
    foo:=len(dfOutput)/6
    renglones:=make([][]string,foo)
    for i:=0; i<foo;i++{
	renglones[i]=dfOutput[i*6:(i*6)+6]
    }
    return renglones
}

func FormaterDiskInfo(foo [][]string)(bar Format_DF){
	listu:=make([]Disk_DF,len(foo))
	for i:=0;i<len(foo);i++{
		listu[i].Filesystem=foo[i][0]
		listu[i].Mem,_=strconv.Atoi(foo[i][1])
		listu[i].Used,_=strconv.Atoi(foo[i][2])
		listu[i].Avaible,_=strconv.Atoi(foo[i][3])
		listu[i].UsePercent=foo[i][4]
		listu[i].Mount=foo[i][5]
	}
	ret:=Format_DF{Title:"Discos",Todos:listu}
	return ret
}

func Mount(disco Disk_DF, MountPoint string)  {
	exec.Command("sudo","mount",disco.Filesystem,MountPoint)
}

func Umount(disco Disk_DF)  {
	exec.Command("sudo","umount",disco.Mount)
}

func GetDisks() (foo Format_lsblk){
	cmd := exec.Command("lsblk", "-J", "-oNAME,SIZE,TYPE,MOUNTPOINTS","-l")
	content, _ := cmd.CombinedOutput()
	var System System_lsblk
	json.Unmarshal(content, &System)

	ret:=Format_lsblk{Title:"TEST",Todos:System}
	return ret
}
