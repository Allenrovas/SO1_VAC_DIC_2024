# Modulos de Kernel

## Herramientas a Utilizar

* __GCC (GNU Compiler Collection)__: Es un conjunto de compiladores creados por el proyecto GNU. Es un software libre y de código abierto. GCC es compatible con varios lenguajes de programación, incluyendo C, C++, Objective-C, Fortran, Ada, Go y D. Originalmente, GCC era un acrónimo para GNU C Compiler (Compilador C de GNU). Sin embargo, GCC ha sido expandido para compilar muchos otros lenguajes además de C.

* __Make__: es una herramienta de construccion utilizada en programación para compilar automáticamente programas y bibliotecas. Es muy utilizada en sistemas Unix y derivados, como Linux. Su función principal es la de automatizar el proceso de compilación de programas, definidas en archivos llamados makefiles.

## Instalación de GCC y Make

Verificar si ya se encuentran instalados:

```sh
gcc --version
make --version
```

Si solo falta make

```sh
sudo apt install make
```

Si falta gcc

```sh
sudo apt install build-essential
sudo apt-get install manpages-dev
```

## Modulos de Kernel

### Compilar un archivo

```sh
make all
```

### Limpiar archivos generados

```sh
make clean
```

### Instalar un modulo

```sh
sudo insmod <<nombre_modulo>>.ko
```

### Obtener los mensajes de entrada y salida del modulo

```sh
sudo dmesg
```

### Verificar si el modulo se encuentra cargado

```sh
cd /proc
```

### Listar los modulos cargados

```sh
ls
```

### Leer el contenido del modulo

```sh
cat /proc/<<nombre_modulo>>
```

### Eliminar un modulo

```sh
sudo rmmod <<nombre_modulo>>.ko
```
