package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	//"os/exec"

	//"github.com/go-ini/ini"
)

//import "github.com/vaughan0/go-ini"

//	GLOBAL
//host allow: Solo computadoras en la red local
//workgourp:
//server string:
//guest account: cuanta de invitado
//max log sizie: tamaño en kb de los archivos de log
//passdb backend:
//interfaces: configuracion de interfaces
//wins support: soporte para windows

//	SHARE DEFINITIONS
// #######todo##########
//hosts allow = 192.168.0.0/16
//hosts deny= 0.0.0.0/0
type Configuration struct{
Value string
Variable string
}

type ConfigurationDefinition struct{
Title string
Contents []Configuration
}

func WriteToFile(Texto string, File string, location int) {
	// Read Write Mode
	file, err := os.OpenFile(File, os.O_RDWR, 0644)
	
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	
	len, err := file.WriteString(Texto)
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

func CreateConfiguration(Configuration ConfigurationDefinition){
	title:="[" +Configuration.Title +"]\n"
	elementsLen:=len(Configuration.Contents)+1
	s:=make([]string,elementsLen)
	s[0]=title
	for i:=0;i<elementsLen-1;i++{
	s[i+1]=Configuration.Contents[i].Variable+" = "+Configuration.Contents[i].Value+"\n"
	}

	for i:=0;i<elementsLen;i++{
	fmt.Printf("%s",s[i])
	}

}


func main()  {
 var test Configuration
 test.Variable="Configureacion1"
 test.Value="Valor"
 var test2 Configuration
 test2.Variable="Configureacion2"
 test2.Value="Valor2"
 var All ConfigurationDefinition
 All.Title="Configuracion Total"
  s := make([]Configuration, 2)
  s[0]=test
  s[1]=test2
 All.Contents=s
 CreateConfiguration(All)
//	file, err := ini.Load("/etc/samba/smb.conf")
//	   if err != nil {
//	        fmt.Printf("Fail to read file: %v", err)
//	        os.Exit(1)
//	    }
//	fmt.Println(file.SectionStrings())
//	file.NewSection("new section")
}
