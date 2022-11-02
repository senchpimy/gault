package main


//
//import (
//	"crypto/sha512"
//	"encoding/base64"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"os"
//	"strings"
//	//"github.com/gorilla/sessions"
//	//"models"
//)
//
////var store= sessions.NewCookieStore([]byte("mysession"))
//
var loginnotworking bool= true
//type User struct{
//User_name string
//Password string
//}
//
//func login(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "GET":
//		http.ServeFile(w,r,"./login.html")
//	case "POST":
//		if err := r.ParseForm(); err !=nil{
//			fmt.Fprintf(w,"ParseForm() err: v%",err)
//			return
//		}
//
//		fmt.Fprintf(w,"Post form website r.postfrom =%v \n",r.PostForm)
//		name:=r.FormValue("name")
//		sha_512:=sha512.New()
//		sha_512.Write([]byte(name))
//		fmt.Fprintf(w,"value = %s\n",name)
//		fmt.Fprintf(w,"value in sha_512 = \t%s",base64.StdEncoding.EncodeToString(sha_512.Sum(nil)))
//		GetPasswordConfirmation("test",name)
//
//
//	default: fmt.Fprintf(w,"Error")
//	}
//}
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////					Recibe un usuario y una contraseña, Lee la base de Usuarios y hace algo cuando la contraseña es correcta o incorrecta
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//func GetPasswordConfirmation(User string,Password string){
//	UsersInGault:= GetGaultUsers()
//	exist_user:=false
//	 for _,item:= range UsersInGault{
//		if item.User_name==User{
//			exist_user=true
//			//password:=sha512.New()
//			//password.Write([]byte(Password))
//			//if item.Password==base64.StdEncoding.EncodeToString(password.Sum(nil)){
//			if item.Password==Password{
//				fmt.Println("LoggedSuccedful")
//				fmt.Println(Password)
//				fmt.Println(item.Password)
//				fmt.Println(item.User_name)
//				//fmt.Println(base64.StdEncoding.EncodeToString(password.Sum(nil)))
//				break
//			}else{
//				fmt.Println("Passwords Don match");
//				fmt.Println(Password)
//				fmt.Println(item.Password)
//				fmt.Println(item.User_name)
//				//fmt.Println(base64.StdEncoding.EncodeToString(password.Sum(nil)))
//				break
//			}
//		}
//	 }
//	 if exist_user==false{fmt.Println("User dont found")}
//}
//
//func GetGaultUsers()(foo []User){  //completo
//    file, err := os.Open("./passwords")
//    if err != nil {
//        log.Fatal(err)
//    }
//    data, err := ioutil.ReadAll(file)
//    if err != nil {
//        log.Fatal(err)
//    }
//    data_string:=string(data)
//
//    data_split:=strings.Split(data_string,"\n")
//    Users:=make([]User,len(data_split)-1)
//    for index,line:= range data_split[:len(data_split)-1]{
//	    data_from_line:=strings.Split(line,":")
//	    Users[index].Password=data_from_line[1]
//	    Users[index].User_name=data_from_line[0]
//
//    }
//    return Users
//}
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//func main(){
////	port := ":3000"
////	mux := http.NewServeMux()
////	mux.HandleFunc("/", login)
////	mux.HandleFunc("/login", login)
////	http.ListenAndServe(port, mux)
//	GetPasswordConfirmation("test","cjeslapolla")
//}
