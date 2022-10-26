package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

var mainpage = "/"
var discospage = "/discosDisponibles"
var discosmontadospage = "/discos"
var sambapage = "/SambaConfi"
///////////////////////////////////////////////////////////////////////////////////////////////////
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		w.Write(readHtmlFromFile("./404.html"))
	}
}

func readHtmlFromFile(fileName string) ([]byte) {

    bs, _ := ioutil.ReadFile(fileName)

    return bs
}

func INIT()  {
	CreateParentDir()
	MountByFile()
}
///////////////////////////////////////////////////////////////////////////////////////////////////
func indexHandler(w http.ResponseWriter, r *http.Request) {
	    if r.URL.Path != mainpage {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	w.Write(readHtmlFromFile("./index.html"))
}
///////////////////////////////////////////////////////////////////////////////////////////////////
func DiscosMontados(w http.ResponseWriter ,r *http.Request)  {
	if r.URL.Path != discosmontadospage {
        	errorHandler(w, r, http.StatusNotFound)
        	return
	}
	switch r.Method {
	case "GET":
		Data:=FormaterDiskInfo(GetInfoSystem())
		t:=template.Must(template.ParseFiles("./discos.html"))
		t.Execute(w,Data)
	case "POST":
		fmt.Println("POST")
		if err := r.ParseForm(); err !=nil{
			fmt.Fprintf(w,"ParseForm() err: v%",err)
			return
		}

		diskUuid:=r.FormValue("diskselected")
		Umount(diskUuid)

	default: fmt.Fprintf(w,"Error")
	}
}
///////////////////////////////////////////////////////////////////////////////////////////////////
func DiscosDisponibles(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != discospage {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	switch r.Method {
	case "GET":
		Data:=GetDisks()
		t:=template.Must(template.ParseFiles("./discosDisponibles.html"))
		t.Execute(w,Data)
	case "POST":
		fmt.Println("POST")
		if err := r.ParseForm(); err !=nil{
			fmt.Fprintf(w,"ParseForm() err: v%",err)
			return
		}

		diskUuid:=r.FormValue("diskselected")
		VerifyDisk(diskUuid)

	default: fmt.Fprintf(w,"Error")
	}
}
///////////////////////////////////////////////////////////////////////////////////////////////////
func SambaConfiguration(w http.ResponseWriter, r *http.Request)  {
	    if r.URL.Path != sambapage {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	switch r.Method {
	case "GET":
		Configuration:=GetAllConfigurations()
		t:=template.Must(template.ParseFiles("./samba.html"))
		t.Execute(w,Configuration)
	case "POST":
		fmt.Println("POST")
		if err := r.ParseForm(); err !=nil{
			fmt.Fprintf(w,"ParseForm() err: v%",err)
			return
		}

		var NewShare Share
		ConfigurationsReceived:=make([]Configuration,12)
		i:=0
		for key, element:= range r.Form{
			if key=="Titulo"{
				NewShare.Title=r.FormValue("Titulo")
				continue
			}
			ConfigurationsReceived[i].Variable=key
			ConfigurationsReceived[i].Value=strings.Join(element," ")
			i++
		}
		NewShare.Contents=ConfigurationsReceived
		VerifyShare(NewShare)

	default: fmt.Fprintf(w,"Error")
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
//         cmd := exec.Command("id", "-u")
//         output, err := cmd.Output()
//
//         if err != nil {
//                 log.Fatal(err)
//         }
//
//
//         // 0 = root, 501 = non-root user
//         i, err := strconv.Atoi(string(output[:len(output)-1]))
//
//         if err != nil {
//                 log.Fatal(err)
//         }
//
//         if i == 0 {
//                 log.Println("root")
//         } else {
//                 log.Fatal("Not root")
//         }
//
	port := ":3000"
	mux := http.NewServeMux()
	mux.HandleFunc(mainpage, indexHandler)
	mux.HandleFunc(discosmontadospage, DiscosMontados)
	mux.HandleFunc(discospage, DiscosDisponibles)
	mux.HandleFunc(sambapage, SambaConfiguration)
	//mux.HandleFunc("/login", login)
	http.ListenAndServe(port, mux)

}
