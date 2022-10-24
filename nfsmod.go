package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/go-ini/ini"
)

//	GLOBAL
//Puertos 20 , 21
//systemctl enable rpcbind nfs-server
//systemctl start rpcbind nfs-server
//systemctl restart rpcbind nfs-server
//systemctl reload rpcbind nfs-server
//systemctl status rpcbind nfs-server
//
// Crear un usuario y asignarlo a una carpeta 
// useradd -g ftp -d /var/www/html/carpetaFTP usuarioftp
// passwd usuarioftp
//
//useradd: Esta es la acción que permite agregar un nuevo usuario en el sistema
//-g: Explica a que grupo va a pertenecer el usuario nuevo que vamos a agregar, en este caso va a pertenecer al grupo ftp.
//-d: Este indica a que directorio va a poder acceder al momento de conectarnos de un cliente FTP, en este caso podrá acceder a /var/www/html/carpetaFTP
//usuarioftp: Es el nombre del usuario que vamos a agregar.
//
//
//	ASIGNAR LA CARPETA A EL USUARIO 
// chown usuarioftp.ftp /var/www/html/carpetaFTP -R



func ExistexotsConf()  { //completado
	if _, err := os.Stat("/etc/exports"); errors.Is(err, os.ErrNotExist) {
		log.Fatal("Samba Config File 'rpcbind nfs-server.conf' does not exist")
		}
}

