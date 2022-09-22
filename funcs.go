package main

import (
	//    "fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

//blkid
//sudo fdisk -l
type Disk struct{
	Filesystem string
	Mem int
	Used int
	Avaible int
	UsePercent string
	Mount string
}

type Format struct{
Title string
Todos []Disk
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

func FormaterDiskInfo(foo [][]string)(bar Format){
	listu:=make([]Disk,len(foo))
	for i:=0;i<len(foo);i++{
		listu[i].Filesystem=foo[i][0]
		listu[i].Mem,_=strconv.Atoi(foo[i][1])
		listu[i].Used,_=strconv.Atoi(foo[i][2])
		listu[i].Avaible,_=strconv.Atoi(foo[i][3])
		listu[i].UsePercent=foo[i][4]
		listu[i].Mount=foo[i][5]
	}
	ret:=Format{Title:"Discos",Todos:listu}
	return ret
}

func Mount(disco Disk, MountPoint string)  {
	exec.Command("sudo","mount",disco.Filesystem,MountPoint)
}

func Umount(disco Disk)  {
	exec.Command("sudo","umount",disco.Mount)
}

