# Docker Compose

### Construir imagenes

```sh
docker compose build
```


### Levantar contenedores

```sh
docker compose up -d
```

NOTA: El flag -d indica que el contenedor se ejecutará en segundo plano.

NOTA: Si aun no se han construido las imagenes se puede agregar el comando --build al final de la instrucción anterior, de la siguiente manera, para construir las imagenes y levantar los contenedores.:

```sh
docker compose up -d --build
```

### Detener contenedores

```sh
docker compose down
```

### Si el nombre del archivo .yml es diferente a docker-compose.yml, utilizaremos el siguiente comando:

```sh
docker compose -f <<nombre_archivo.yml>> up -d
docker compose -f <<nombre_archivo.yml>> down
```

# Docker Volumes

### Listar Volumenes

```sh
docker volume ls
```

### Crear un Volumen

```sh
docker volume create <<nombre_volumen>>
```

### Eliminar un Volumen

```sh
docker volume rm <<nombre_volumen>>
```

### Eliminar todos los Volumenes

```sh
docker volume prune
```

# MYSQL

### Descargar imagen de MySQL

```sh
docker pull mysql
```

### Referencia: [MYSQL](https://hub.docker.com/_/mysql)

### Crear un contenedor de MySQL

```sh
docker run --name <<nombre_contenedor>> -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=<<password>> mysql
```

Ejemplo:

```sh
docker run --name mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
```

## Acceder al contenedor de MySQL

```sh
docker exec -it <<nombre_contenedor>> mysql -u root -p<<password>>
```

Ejemplo:

```sh
docker exec -it mysql mysql -u root -p123456
```

Aqui se puede ejecutar cualquier comando de MySQL.

## Crear una base de datos

```sh
CREATE DATABASE <<nombre_base_de_datos>>;
```

## Usar una base de datos

```sh
USE <<nombre_base_de_datos>>;
```


### Con persistencia de datos

```sh
docker run --name <<nombre_contenedor>> -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=<<password>> -mount src=<<nombre_volumen>>,dst=/var/lib/mysql mysql
```

Ejemplo:

```sh
docker run --name mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -mount src=mysqldata,dst=/var/lib/mysql mysql
```

# Workbench

## Instalar Workbench en Ubuntu

### Paso 1: Descargar la version que deseamos en [Workbench](https://dev.mysql.com/downloads/workbench/)


### Paso 2: Instalar dependencias, ubicarse en la carpeta donde se descargo el archivo y ejecutar el siguiente comando:

```sh
sudo dpkg -i <<nombre_archivo.deb>>
```

### Paso extra: si al terminar el proceso anterior existen algunos errores en la instalacion por dependencias, ejecutar el siguiente comando:

```sh
sud apt --fix-broken install
```

# DBeaver

## Instalar DBeaver en Ubuntu

```sh
sudo snap install dbeaver-ce
```