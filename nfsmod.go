package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

//	GLOBAL
//Puertos 20 , 21
//systemctl enable rpcbind nfs-server
//systemctl start rpcbind nfs-server
//systemctl restart rpcbind nfs-server
//systemctl reload rpcbind nfs-server
//systemctl status rpcbind nfs-server
//
type nfsFiles struct{
Location string
Host string
}
type Nfs struct{
	Shares []nfsFiles
	Mounted []Disk_DF
}

func ListExports() (foo Nfs) { //completado
	//cmd:=exec.Command("exportfs","-v")
	cmd:=exec.Command("cat","./exports")
    	out, _ := cmd.CombinedOutput()
	f:=strings.Fields(string(out))
	fo:=make([]nfsFiles, len(f)/2)
	j:=0
	for i:=1; i<len(f) ;i+=2{
		fo[j].Host=f[i]
		fo[j].Location=f[i-1]
        j++

	}
	Data:=FormaterDiskInfo(GetInfoSystem())
	bar:=Nfs{Shares:fo, Mounted:Data.Todos}
	return bar
}

func ExistexotsConf()  { //completado
	if _, err := os.Stat("/etc/exports"); errors.Is(err, os.ErrNotExist) {
		log.Fatal("Samba Config File 'rpcbind nfs-server.conf' does not exist")
		}
}

func CreateExport(path, permissions, red, options string){
	// Read Write Mode
	file, err := os.OpenFile("./exports", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	
	len, err := file.WriteString(path+" "+red+"("+permissions+","+options+")\n")
	if err != nil {
		log.Fatalf("failed writing to file: %s %s", err,len)
	}
}

func DeleteNfs(nfs string)  {
	data, err := ioutil.ReadFile("./exports")
	if err != nil {
	    log.Fatal(err)
	}
	file := string(data)
	temp := strings.Split(file, "\n")

        newFile, err := os.Create("./exports.after")
	if err != nil {
        log.Fatal(err)
    	}
    	defer newFile.Close()
	temp=temp[:len(temp)-1]

	for _, item := range temp {
		splited:=strings.Fields(item)
		if len(splited)>0{
		if nfs==splited[0] {
			continue
		}}
		newFile.Write([]byte(item+"\n"))
	}

	err = exec.Command( "mv","./exports","./exports.bak").Run()
	if err != nil {log.Fatal(err)}
	err = exec.Command( "mv","./exports.after","./exports").Run()
	if err != nil {log.Fatal(err)}
}
