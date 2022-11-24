package main

import (
    "fmt"
    "net/http"
"encoding/json"
)

type InputFromPage struct {
    Input string
}

type Response struct {
    Output string
}

func HandleButtons(w http.ResponseWriter, r *http.Request){
    tpl.ExecuteTemplate(w,"404.html",nil)
    decoder := json.NewDecoder(r.Body)
    var t InputFromPage
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    switch t.Input{
	case "logout":
		logoutHandler(w,r)
		ServerResponse:= Response{Output:"Done"}
		byteArray, err := json.Marshal(ServerResponse)
		w.Write(byteArray)
        	if err != nil {fmt.Println(err)}
		w.Write(byteArray)

	case "vsftpd.service":
		RestartFtp()

	case "sshd.service":
		//Restartssh()

	case "nfs-server.service":
//		RestartNfs()

	case "smb.service":
		RestartSamba()

	default:
		fmt.Println("Input no encontrado")
		fmt.Println(t.Input)
    }
}


