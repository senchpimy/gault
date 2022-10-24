package main

import (
	"errors"
	"log"
	"os"
)

//	GLOBAL
//Puertos 20 , 21
//systemctl enable vsftpd.service
//systemctl start vsftpd
//systemctl restart vsftpd
//systemctl reload vsftpd
//systemctl status vsftpd
//systemctl enable vsftpd
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

