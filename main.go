package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	//"github.com/gorilla/sessions"
	"github.com/gorilla/context"
)

var mainpage = "/"
var discospage = "/discosDisponibles"
var discosmontadospage = "/discos"
var sambapage = "/SambaConfi"
var UserConfig = "/UserConfig"
var Login = "/login"

var tpl *template.Template
///////////////////////////////////////////////////////////////////////////////////////////////////
func errorHandler(w http.ResponseWriter, r *http.Request, PageName string) (foo bool){
	status:=http.StatusNotFound
	if r.URL.Path != PageName {
		w.WriteHeader(status)
		if status == http.StatusNotFound {
			w.Write(readHtmlFromFile("./404.html"))
	        return true
		}
	}

	session, _ := storeOfsessions.Get(r, "session")
	_, ok := session.Values["userID"]
	if !ok {
 		http.Redirect(w, r, Login, http.StatusFound) // http.StatusFound is 302
        	return true
	}else{return false}
}

func readHtmlFromFile(fileName string) ([]byte) {
    bs, _ := ioutil.ReadFile(fileName)
    return bs
}

func INIT()  {
	CreateParentDir()
	MountByFile()
	fmt.Println("INIT pasado")
}
///////////////////////////////////////////////////////////////////////////////////////////////////
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if errorHandler(w,r,mainpage) {
		return
	}
	tpl.ExecuteTemplate(w,"index.html",nil)
}
///////////////////////////////////////////////////////////////////////////////////////////////////
func DiscosMontados(w http.ResponseWriter ,r *http.Request)  {
	if errorHandler(w,r,discosmontadospage){
		return
	}

	switch r.Method {
	case "GET":
		tpl.ExecuteTemplate(w, "discos.html", FormaterDiskInfo(GetInfoSystem()))
	case "POST":
		fmt.Println("POST")
		if err := r.ParseForm(); err !=nil{
			fmt.Fprintf(w,"ParseForm() err: v%",err)
			return
		}

		diskUuid:=r.FormValue("diskselected")
		Umount(diskUuid)
		Data:=FormaterDiskInfo(GetInfoSystem())
		tpl.ExecuteTemplate(w, "discos.html", Data)

	default: fmt.Fprintf(w,"Error")
	}
}
///////////////////////////////////////////////////////////////////////////////////////////////////
func DiscosDisponibles(w http.ResponseWriter, r *http.Request)  {
	if errorHandler(w, r, discospage){
		return
    }
	switch r.Method {
	case "GET":
		tpl.ExecuteTemplate(w, "discosDisponibles.html", GetDisks())
	case "POST":
		fmt.Println("POST")
		if err := r.ParseForm(); err !=nil{
			fmt.Fprintf(w,"ParseForm() err: v%",err)
			return
		}

		diskUuid:=r.FormValue("diskselected")
		VerifyDisk(diskUuid)
		Data:=GetDisks()
		t:=template.Must(template.ParseFiles("./discosDisponibles.html"))
		t.Execute(w,Data)

	default: fmt.Fprintf(w,"Error")
	}
}
///////////////////////////////////////////////////////////////////////////////////////////////////
func SambaConfiguration(w http.ResponseWriter, r *http.Request)  {
    if errorHandler(w,r,sambapage){
	return
    }
	switch r.Method {
	case "GET":
		tpl.ExecuteTemplate(w, "samba.html", GetAllConfigurations())
	case "POST":
		if err := r.ParseForm(); err !=nil{
			fmt.Fprintf(w,"ParseForm() err: v%",err)
			return
		}

		if len(r.Form)!=1{
			Correct:=true
			var NewShare Share
			ConfigurationsReceived:=make([]Configuration,12)
			i:=0
			for key, element:= range r.Form{
				if key=="Titulo"{
					NewShare.Title=r.FormValue("Titulo")
					continue
				}
				if key=="Delete"{continue}
				if key=="valid users"{
					userslist, valid:=UsersExist(strings.Join(element," "))
					fmt.Println(valid)
					if valid==true{
						ConfigurationsReceived[i].Variable=key
						UsersFormated:=strings.Join(element," ")
						ConfigurationsReceived[i].Value=strings.Replace(UsersFormated,","," ",-1)
						i++
						continue
					}else{
						CreateError("The following users don exist:")
						for _, item := range userslist{
							CreateError(item)

						}
						CreateError("Imposible de Crear El Share")
						Correct=false
						break
					}
				}
				ConfigurationsReceived[i].Variable=key
				ConfigurationsReceived[i].Value=strings.Join(element," ")
				i++
			}
			NewShare.Contents=ConfigurationsReceived
			if Correct==true{VerifyShare(NewShare)}else{CreateError("Un Error Sucedio Imposible de Crear Share")}
		}else{
		Share:=r.FormValue("Delete")
		DeleteShare(Share)
		}
		tpl.ExecuteTemplate(w, "samba.html", GetAllConfigurations())

	default: fmt.Fprintf(w,"Error")
	}
}
///////////////////////////////////////////////////////////////////////////////////////////////////
func Users(w http.ResponseWriter, r *http.Request)  {
    if errorHandler(w,r,UserConfig){
	return
    }
	switch r.Method {
	case "GET":
		tpl.ExecuteTemplate(w, "users.html", GetUsers())
	case "POST":
		fmt.Println("POST")
		if err := r.ParseForm(); err !=nil{
			fmt.Fprintf(w,"ParseForm() err: v%",err)
			return
		}

		User:=r.FormValue("User")
		Passw1:=r.FormValue("Passw1")
		Passw2:=r.FormValue("Passw2")
		TypeOfUser:=r.FormValue("Type")
		AddUser(User,Passw1,Passw2,TypeOfUser)
		tpl.ExecuteTemplate(w, "users.html", GetUsers())

	default: fmt.Fprintf(w,"Error")
	}
}
///////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	INIT()
	port := ":3000"
	tpl, _ = template.ParseGlob("templates/*.html")
	//mux := http.NewServeMux()
	//mux.HandleFunc(mainpage, indexHandler)
	//mux.HandleFunc(discosmontadospage, DiscosMontados)
	//mux.HandleFunc(discospage, DiscosDisponibles)
	//mux.HandleFunc(sambapage, SambaConfiguration)
	//mux.HandleFunc(ftpPage, FTPConfiguration)
	//mux.HandleFunc(UserConfig, Users)
	////mux.HandleFunc("/login", login)
	//http.ListenAndServe(port, mux)

	http.HandleFunc(mainpage, indexHandler)
	http.HandleFunc(discosmontadospage, DiscosMontados)
	http.HandleFunc(discospage, DiscosDisponibles)
	http.HandleFunc(sambapage, SambaConfiguration)
	//http.HandleFunc(ftpPage, FTPConfiguration)
	http.HandleFunc(UserConfig, Users)
	http.HandleFunc(Login, login)
	http.ListenAndServe(port, context.ClearHandler(http.DefaultServeMux))

}
