package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

//	GLOBAL
//Puertos 20 , 21
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



func ExistFTPConf()  { //completado
	if _, err := os.Stat("/etc/vsftpd/vsftpd.conf"); errors.Is(err, os.ErrNotExist) {
		log.Fatal("Samba Config File 'vsftpd.conf' does not exist")
		}
}

 func StartFtp(){
	
    cmd := exec.Command("systemctl", "start", "vsftpd")
    err := cmd.Run()
    if err != nil {log.Fatal(err)}
 }

 func EnableFtp(){
	
    cmd2 := exec.Command("systemctl", "enable", "vsftpd")
    err2 := cmd2.Run()
    if err2 != nil {log.Fatal(err2)}
 }

 func RestartFtp(){
	
    cmd2 := exec.Command("systemctl", "restart", "vsftpd")
    err2 := cmd2.Run()
    if err2 != nil {log.Fatal(err2)}
 }

 func ReloadFtp(){
	
    cmd2 := exec.Command("systemctl", "reload", "vsftpd")
    err2 := cmd2.Run()
    if err2 != nil {log.Fatal(err2)}
 }

// func StatuFtp(){
//	
//    cmd2 := exec.Command("systemctl", "status", "vsftpd")
//    err2 := cmd2.Run()
//    if err2 != nil {log.Fatal(err2)}
// }

