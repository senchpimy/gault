# GAULT

![Badge](https://img.shields.io/badge/License-MIT-blue)
![Badge](https://img.shields.io/badge/Version-1.0.0-green)
![Badge](https://img.shields.io/badge/Go-1.21-blue?logo=go)

**GAULT** es una alternativa ligera a OpenMediaVault escrita completamente en Go.
Permite la administración de discos, usuarios del sistema, y configuraciones de Samba y NFS desde una interfaz web moderna, sin necesidad de depender de entornos pesados.

---

## ✨ Características

* 📊 Dashboard con uso de CPU y discos en tiempo real
* 🗂 Gestión de discos montados y disponibles
* 🧑‍💻 Creación y gestión de usuarios del sistema
* 🔐 Sistema de login y logout
* 📁 Interfaz para configurar recursos compartidos vía **Samba**
* 🌐 Soporte para exportaciones vía **NFS**
* ⚙️ Interfaz basada en plantillas HTML (`html/template`)
* 🔒 Protección por sesión: todas las rutas verifican autenticación
* 📄 Registro de errores en archivo de log

---

## ⚙️ Requisitos

Este programa requiere ejecutarse como **root**.

Además, necesita que los siguientes programas estén instalados en el sistema:

* `lsblk`
* `df`
* `samba`
* `smbclient`
* `useradd`

---

## 🚀 Cómo ejecutar

```bash
sudo go run main.go
```

El servidor web se ejecuta por defecto en:

```
http://localhost:3000
```

Debes tener disponible una carpeta `templates/` con las vistas `.html` necesarias y la hoja de estilos en `templates/styles/style.css`.

---

## 🔧 En desarrollo (To Do)

* Soporte para **FTP**
* Recarga y activación de **Samba** tras cambios
* Visualización de mensajes de error desde la página web
* Control del **firewall** desde la interfaz
* Login 

---

## 📄 Licencia

Este proyecto está bajo la licencia [MIT](https://opensource.org/licenses/MIT).

---
