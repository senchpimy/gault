package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"html/template"
)

var mainpage = "/"
var discospage = "/discos"
var discosmontadospage = "/discosDisponibles"
var sambapage = "/SambaConfi"

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
        fmt.Fprint(w, Logged)
    }
}

func readHtmlFromFile(fileName string) ([]byte) {

    bs, _ := ioutil.ReadFile(fileName)

    return bs
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	    if r.URL.Path != mainpage {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	w.Write(readHtmlFromFile("./index.html"))
}

func DiscosMontados(w http.ResponseWriter ,r *http.Request)  {
	    if r.URL.Path != discosmontadospage {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	Data:=FormaterDiskInfo(GetInfoSystem())
	t:=template.Must(template.ParseFiles("./discos.html"))
	t.Execute(w,Data)
}

func DiscosDisponibles(w http.ResponseWriter, r *http.Request)  {
	    if r.URL.Path != discospage {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	Data:=GetDisks()
	t:=template.Must(template.ParseFiles("./discosDisponibles.html"))
	t.Execute(w,Data)
}

func SambaConfiguration(w http.ResponseWriter, r *http.Request)  {
	    if r.URL.Path != sambapage {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	Configuration:=GetAllConfigurations()
	t:=template.Must(template.ParseFiles("./samba.html"))
	t.Execute(w,Configuration)
}

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
