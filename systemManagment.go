package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
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
	_,err:=exec.Command( "mount","/dev/"+disco,MountPoint).Output()
	fmt.Println(err)
}

func Umount(disco string)  {
	_,err:=exec.Command("umount", "/dev/"+disco).Output()
	fmt.Println(err)
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

func CreateMountDir(dirname string)(foo string){
	md5:=md5.New()
	io.WriteString(md5,dirname)
	dirname=base64.StdEncoding.EncodeToString(md5.Sum(nil))
	cmd := exec.Command("mkdir", "/run/gault/"+dirname, "-p")
	cmd.CombinedOutput()
	return dirname
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
	if _, err := f.Write([]byte(disk+":"+MountPoint)); err != nil {
	    log.Fatal(err)
	    }
	if err := f.Close(); err != nil {
	    log.Fatal(err)
	}

}

func VerifyDisk(disk string)  { //Recibe UUID del disco
	dirname:=CreateMountDir(disk)     //Crea una carpeta en donde se va a montar...
	if dirname in archivo{            //...esta carpeta es unica a la uuid del disco
	//si el directorio ya existe y el disco ya esta en el archivo de configuracion entonces solo montar	
	}else{
		//si no agregar a la configuracion y montar
		AddDiskToConfig(disk,dirname)
		Mount()
	}
	
}

//func main(){
//}
