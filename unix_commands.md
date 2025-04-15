## 🐧 Comandos Básicos de Unix/Linux/Shell

Comandos esenciales para navegar, gestionar archivos y procesos en sistemas tipo Unix.

| Comando | Descripción | Ejemplo |
|---------|-------------|---------|
| **`clear`** | para limpiar la pantalla en la terminal| `clear` |
| **`pwd`** | Muestra la ruta absoluta del directorio actual. | `pwd` → `/home/usuario/proyectos` |
| **`ls`** | Lista archivos y directorios. | `ls -l` (lista detallada) |
| **`cd`** | Cambia de directorio. | `cd /ruta` o `cd ..` (subir un nivel) |
| **`mkdir`** | Crea un nuevo directorio. | `mkdir nombre_directorio` |
| **`rm`** | Elimina archivos/directorios. | `rm archivo.txt` (¡Cuidado! No va a la papelera) |
| **`cp`** | Copia archivos o directorios. | `cp origen destino` |
| **`mv`** | Mueve o renombra archivos/directorios. | `mv viejo.txt nuevo.txt` |
| **`cat`** | Muestra el contenido de un archivo. | `cat archivo.txt` |
| **`grep`** | Busca patrones en archivos. | `grep "texto" archivo.log` |
| **`chmod`** | Cambia permisos de archivos. | `chmod +x script.sh` (da permisos de ejecución) |
| **`ps`** | Lista procesos en ejecución. | `ps aux` |
| **`kill`** | Termina un proceso por su PID. | `kill -9 1234` |
| **`top`** | Monitoriza procesos en tiempo real. | `top` |
| **`df`** | Muestra espacio en disco. | `df -h` (en formato legible) |
| **`du`** | Muestra uso de espacio por archivos. | `du -sh *` (resumen por carpeta) |
| **`tar`** | Comprime/descomprime archivos. | `tar -czvf archivo.tar.gz carpeta/` |
| **`ssh`** | Conexión remota segura. | `ssh usuario@servidor` |
| **`scp`** | Copia archivos entre máquinas (vía SSH). | `scp archivo.txt usuario@servidor:/ruta` |
| **`echo`** | Imprime texto o variables. | `echo $PATH` |
| **`man`** | Manual de un comando. | `man ls` |

### 📌 Tips:
- Usa `comando --help` para ver ayuda rápida (ej: `ls --help`).  
- Presiona **`Ctrl + C`** para cancelar un proceso en ejecución.  
- Usa **`Tab`** para autocompletar nombres de archivos/directorios.  

> ℹ️ **Nota**: Estos comandos funcionan en terminales de Linux, macOS, WSL (Windows) y Git Bash.  
> Para usarlos en Windows CMD, muchos requieren equivalencias (ej: `dir` en vez de `ls`).