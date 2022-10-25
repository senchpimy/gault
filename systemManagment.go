package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

//sudo fdisk -l: mucha info

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

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

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func FormaterDiskInfo(foo [][]string)(bar Format_DF){
	listu:=make([]Disk_DF,len(foo))
	for i:=0;i<len(foo);i++{
		if foo[i][0]=="tmpfs" || foo[i][0]=="dev" || foo[i][0]=="run" || foo[i][5]=="/"{
			continue
		}else{
		listu[i].Filesystem=foo[i][0]
		listu[i].Mem,_=strconv.Atoi(foo[i][1])
		listu[i].Used,_=strconv.Atoi(foo[i][2])
		listu[i].Avaible,_=strconv.Atoi(foo[i][3])
		listu[i].UsePercent=foo[i][4]
		listu[i].Mount=foo[i][5]
		}
	}
	ret:=Format_DF{Title:"Discos",Todos:listu}
	return ret
}

func Mount(disco string, MountPoint string)  {
	_,err:=exec.Command( "mount","/dev/"+disco,MountPoint).Output()
	if err != nil {CreateError(string(err.Error()))}
}

func Umount(disco string)  {
	_,err:=exec.Command("umount", disco).Output()
	if err != nil {CreateError(string(err.Error()))}
}

func GetDisks() (foo System_lsblk){
	cmd := exec.Command("lsblk", "-J", "-oNAME,SIZE,TYPE,MOUNTPOINTS,RM,UUID","-l")
	content, _ := cmd.CombinedOutput()
	var System System_lsblk
	json.Unmarshal(content, &System)
	return System
}

func CreateParentDir(){
	cmd := exec.Command("mkdir", "/run/media/gault", "-p")
	cmd.CombinedOutput()
}

func CreateMountDir() string{
	dirname:=randSeq(10)
	_, err := exec.Command("mkdir", "/run/media/gault/"+dirname, "-p").Output()
	fmt.Println(err)
	return "/run/media/gault/"+dirname
}

func MountByFile(){
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
	        Disks[index].MountPoint=data_from_line[1]
	        Disks[index].Uuid=data_from_line[0]
	
	}
	CurrentDisksConnected:=GetDisks()
	for _,disk:= range CurrentDisksConnected.Blockdevices{
		for _,disksConfigured:= range Disks{
			if disk.Uuid == disksConfigured.Uuid{
				Mount(disk.Name,disksConfigured.MountPoint)
			}
		}
	}
}

func MountByUUID(uuid string, dir string){
	a:=GetDisks()	
	for _,item := range a.Blockdevices{
		if item.Uuid==uuid{
			Mount(item.Name, dir)
		}
	}
}

func AddDiskToConfig(disk string, MountPoint string){
	f, err := os.OpenFile("./disks", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
	    log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.Write([]byte(disk+":"+MountPoint+"\n")); err != nil {
	    log.Fatal(err)
	    }
	if err := f.Close(); err != nil {
	    log.Fatal(err)
	}

}


func VerifyDisk(diskUuid string)  { //Recibe UUID del disco
	if diskUuid!="null"{
		dirname:=CreateMountDir()         //Crea una carpeta en donde se va a montar...
		by,_:=ioutil.ReadFile("./disks")  //...esta carpeta es unica a la uuid del disco
		file:=string(by)            
		fmt.Println(string(by))
		if strings.Contains(file, diskUuid){
		//si el directorio ya existe y el disco ya esta en el archivo de configuracion entonces solo montar
			MountByFile()
		}else{
			//si no agregar a la configuracion y montar
			AddDiskToConfig(diskUuid,dirname)
			MountByUUID(diskUuid,dirname)
		}
	}
}
//func main(){
//}
