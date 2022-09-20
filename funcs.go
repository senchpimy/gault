package main

import (
    "fmt"
    "log"
    "strings"
    "os/exec"
    "strconv"
)
//blkid
//sudo fdisk -l
type disk struct{
	Filesystem string
	Mem int
	Used int
	Avaible int
	UsePercent string
	Mount string
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

func FormaterDiskInfo(foo [][]string)(list []disk){
	listu:=make([]disk,len(foo))
	for i:=0;i<len(foo);i++{
		listu[i].Filesystem=foo[i][0]
		listu[i].Mem,_=strconv.Atoi(foo[i][1])
		listu[i].Used,_=strconv.Atoi(foo[i][2])
		listu[i].Avaible,_=strconv.Atoi(foo[i][3])
		listu[i].UsePercent=foo[i][4]
		listu[i].Mount=foo[i][5]
	}
	return listu
}
func Mount(disco disk, MountPoint string)  {
	os.exec("sudo","mount",disco.Filesystem,MountPoint)
}
func main()  {
}
