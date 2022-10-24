package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"encoding/json"
	"io/ioutil"
	"os"
)

//sudo fdisk -l: mucha info

type Disk_DF struct{
	Filesystem string
	Mem int
	Used int
	Avaible int
	UsePercent string
	Mount string
}

type Disk struct{
	Uuid string
	MountPoint string
}

type Format_DF struct{
Title string
Todos []Disk_DF
}

type Disk_lbslk struct{
	Name string 
	Size string 
	Type string
	Rm bool
	Uuid string
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

func Mount(disco string, MountPoint string)  {
	cmd:=exec.Command("sudo","mount",disco,MountPoint)
	cmd.CombinedOutput()
}

func Umount(disco string)  {
	cmd:=exec.Command("sudo","umount",disco)
	cmd.CombinedOutput()
}

func GetDisks() (foo System_lsblk){
	cmd := exec.Command("lsblk", "-J", "-oNAME,SIZE,TYPE,MOUNTPOINTS,RM,UUID","-l")
	content, _ := cmd.CombinedOutput()
	var System System_lsblk
	json.Unmarshal(content, &System)
	return System
}

func CreateParentDir(){
	cmd := exec.Command("mkdir", "/run/gault", "-p")
	cmd.CombinedOutput()
}

func GetUUIDandMount()(foo []Disk){
    file, err := os.Open("./disks")
    if err != nil {
        log.Fatal(err)
    }
    data, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }
    data_string:=string(data)

    data_split:=strings.Split(data_string,"\n")
    Disks:=make([]Disk,len(data_split)-1)
    for index,line:= range data_split[:len(data_split)-1]{
	    data_from_line:=strings.Split(line,":")
	    Disks[index].Uuid=data_from_line[1]
	    Disks[index].MountPoint=data_from_line[0]

    }
    return Disks
    }
