package main

import (
	"log"
	"os/exec"
	"strings"
)

type UsersInSystem struct{
Users []string
}

func GetUsers()(foo UsersInSystem){
	out,err:=exec.Command("ls","/home").Output()
	if err != nil{CreateError(err.Error())}
	NumeOfUsers:=strings.Split(string(out),"\n")
	list:=make([]string,len(NumeOfUsers)-1)
	for index,item:= range NumeOfUsers[:len(NumeOfUsers)-1]{
		list[index]=item
	}
	var Test UsersInSystem
	Test.Users=list
	return Test
}

func AddUser(user string,password1 string, password2 string)  {
	if password1==password2{
		password:=password1
		err:=exec.Command("sh","./CreateUser.sh",user,password)
		if err != nil{log.Fatal(err)}
	}else{
		CreateError("Las Contraseñas no son Iguales")
	}
}

func UsersExist(Users string) ([]string, bool){
	Users=strings.Replace(Users," ", "",-1)
	ListOfUsers:=strings.Split(Users,",")
	UsersInSystemSlice:=GetUsers()
	mb := make(map[string]struct{}, len(UsersInSystemSlice.Users))
	for _, x := range UsersInSystemSlice.Users{
	    mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range ListOfUsers {
	    if _, found := mb[x]; !found {diff = append(diff, x)}
	}
	
	if len(diff)==0{
	
	    return diff, true
	}else{
	
	    return diff, false
	}
}

