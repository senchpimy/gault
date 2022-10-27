package main

import (
	"fmt"
	//"os"
	"os/exec"
	"strings"
)

type UsersInSystem struct{
Users []string
}
func GetUsers()(foo UsersInSystem){
	out,err:=exec.Command("ls","/home").Output()
	if err != nil{fmt.Errorf(err.Error())}
	list:=make([]string,len(strings.Split(string(out),"\n")))
	for index,item:= range strings.Split(string(out),"\n"){
		list[index]=item
	}
	var Test UsersInSystem
	Test.Users=list
	return Test
}

func AddUser(user string,password1 string, password2 string)  {
	if password1==password2{
	password:=password1
	}else{
	CreateError("Las Contraseñas no son Iguales")
	}
	err:=exec.Command("sh","./CreateUser.sh",user,password)
	if err != nil{CreateError(err.Error())}
}

//func main(){
//GetUsers()
//}
