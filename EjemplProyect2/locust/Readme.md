# Pasos para el generador de tráfico con Locust

## 1. Crear un archivo .py para generar el json con el que se hará la prueba
```python
import json
import random

# Funciones para generar datos aleatorios
def generar_facultad():
    facultades = ["Ingenieria", "Medicina", "Derecho", "Arquitectura", "Economia", "Veterinaria", "Odontologia"]
    return random.choice(facultades)

def generar_curso():
    cursos = list({"SO1", "LFP", "BD1", "SA", "AYD1", "SO2", "BD2", "AYD2", "IA1"})
    return random.choice(cursos)

def generar_carrera():
    carreras = ["Sistemas", "Industrial", "Civil", "Mecanica", "Electronica", "Quimica", "Biomedica"]
    return random.choice(carreras)

def generar_region():
    regiones = ["METROPOLITANA", "NORTE", "NORORIENTAL", "SURORIENTAL", "CENTRAL", "SUROCCIDENTAL", "NOROCCIDENTAL", "PETEN"]
    return random.choice(regiones)

def generar_json():
    return {
        "curso": generar_curso(),
        "carrera": generar_carrera(),
        "facultad": generar_facultad(),
        "region": generar_region()
    }

def generar_archivo_json(cantidad, nombre_archivo):
    datos = [generar_json() for _ in range(cantidad)]
    with open(nombre_archivo, "w") as file:
        json.dump(datos, file, indent=4)
    print(f"Archivo generado: {nombre_archivo}")

# Bloque principal
if __name__ == "__main__":
    print("Generador de datos para asignación de cursos")
    while True:
        try:
            cantidad = int(input("Ingrese la cantidad de registros a generar: "))
            if cantidad > 0:
                break
            print("Por favor, ingrese un número mayor que 0.")
        except ValueError:
            print("Entrada inválida. Por favor, ingrese un número.")

    nombre_archivo = input("Ingrese el nombre del archivo (sin extensión): ").strip() + ".json"
    print(f"Generando {cantidad} registros...")
    generar_archivo_json(cantidad, nombre_archivo)
```

## 2. Instalar oras para subir el archivo .json como OCI Object Storage
```bash
VERSION="1.2.0"
curl -LO "https://github.com/oras-project/oras/releases/download/v${VERSION}/oras_${VERSION}_linux_amd64.tar.gz"
mkdir -p oras-install/
tar -zxf oras_${VERSION}_*.tar.gz -C oras-install/
sudo mv oras-install/oras /usr/local/bin/
rm -rf oras_${VERSION}_*.tar.gz oras-install/
```

### 2.1 Verificar la instalación 
```bash
oras version
```

## 3. Subir el archivo .json a OCI Object Storage (ya debe estar configurado Harbor)

### 3.1. Iniciar sesión en el registro de contenedores

```bash
oras login --insecure registry.home-k8s.lab -u admin -p <tu-contraseña>
```

### 3.2. Subir el archivo .json a OCI Object Storage

```bash
oras push --insecure registry.home-k8s.lab/project/asignaciones:latest \
    --artifact-type application/json \
    <<nombre-archivo.json>>
```

### 3.3. También se puede subir el archivo .json desde el mismo script de Python

```python
import json
import random
import os

# Funciones para generar datos aleatorios
def generar_facultad():
    facultades = ["Ingenieria", "Medicina", "Derecho", "Arquitectura", "Economia", "Veterinaria", "Odontologia"]
    return random.choice(facultades)

def generar_curso():
    cursos = list({"SO1", "LFP", "BD1", "SA", "AYD1", "SO2", "BD2", "AYD2", "IA1"})
    return random.choice(cursos)

def generar_carrera():
    carreras = ["Sistemas", "Industrial", "Civil", "Mecanica", "Electronica", "Quimica", "Biomedica"]
    return random.choice(carreras)

def generar_region():
    regiones = ["METROPOLITANA", "NORTE", "NORORIENTAL", "SURORIENTAL", "CENTRAL", "SUROCCIDENTAL", "NOROCCIDENTAL", "PETEN"]
    return random.choice(regiones)

def generar_json():
    return {
        "curso": generar_curso(),
        "carrera": generar_carrera(),
        "facultad": generar_facultad(),
        "region": generar_region()
    }

def generar_archivo_json(cantidad, nombre_archivo):
    datos = [generar_json() for _ in range(cantidad)]
    with open(nombre_archivo, "w") as file:
        json.dump(datos, file, indent=4)
    print(f"Archivo generado: {nombre_archivo}")

def hacer_push_oras(nombre_archivo, registro, proyecto, etiqueta):
    # Comando para hacer el push con oras
    comando = f"oras push --insecure {registro}/{proyecto}/{etiqueta} --artifact-type application/json {nombre_archivo}"
    print(f"Ejecutando: {comando}")
    resultado = os.system(comando)
    if resultado == 0:
        print(f"Archivo {nombre_archivo} subido exitosamente al registro {registro}/{proyecto}/{etiqueta}")
    else:
        print(f"Error al subir el archivo {nombre_archivo} al registro {registro}/{proyecto}/{etiqueta}")

# Bloque principal
if __name__ == "__main__":
    print("Generador de datos para asignación de cursos")
    while True:
        try:
            cantidad = int(input("Ingrese la cantidad de registros a generar: "))
            if cantidad > 0:
                break
            print("Por favor, ingrese un número mayor que 0.")
        except ValueError:
            print("Entrada inválida. Por favor, ingrese un número.")

    nombre_archivo = input("Ingrese el nombre del archivo (sin extensión): ").strip() + ".json"
    print(f"Generando {cantidad} registros...")
    generar_archivo_json(cantidad, nombre_archivo)

    # Configuración del registro OCI
    registro = input("Ingrese el dominio del registro OCI (ej. registry.home-k8s.lab): ").strip()
    proyecto = input("Ingrese el nombre del proyecto en Harbor: ").strip()
    etiqueta = input("Ingrese la etiqueta para el artefacto (ej. asignaciones:latest): ").strip()

    # Hacer push al registro OCI
    hacer_push_oras(nombre_archivo, registro, proyecto, etiqueta)
```

## 4. Crear un archivo .py para ejecutar la prueba de carga con Locust y hacer el pull del archivo .json
```python
from locust import HttpUser, between, task
from random import randrange
import json
import os

class ReadFile():
    def __init__(self):
        self.calificaciones = []

    def getData(self):
        size = len(self.calificaciones)
        if size > 0:
            index = randrange(0, size - 1) if size > 1 else 0
            return self.calificaciones.pop(index)
        else:
            print("No hay más datos")
            return None

    def loadFile(self, file_name):
        """
        Carga el archivo JSON desde el nombre proporcionado.
        Si es un directorio, busca el archivo JSON dentro de él.
        """
        try:
            # Verificar si el nombre proporcionado es un directorio
            if os.path.isdir(file_name):
                # Buscar el archivo JSON dentro del directorio
                files = os.listdir(file_name)
                json_file = next((f for f in files if f.endswith(".json")), None)
                if json_file:
                    file_name = os.path.join(file_name, json_file)
                else:
                    raise Exception(f"No se encontró un archivo JSON en el directorio {file_name}")

            # Cargar el archivo JSON
            with open(file_name, "r", encoding="utf-8") as file:
                self.calificaciones = json.loads(file.read())
            print(f"Archivo {file_name} cargado exitosamente.")
        except Exception as e:
            print(f"Error al cargar el archivo: {e}")
            raise

    def pullFromOCI(self, registry, project, artifact, output_file):
        """
        Realiza el pull del archivo JSON desde un registro OCI utilizando oras.
        """
        command = f"oras pull --insecure {registry}/{project}/{artifact} --output {output_file}"
        print(f"Ejecutando: {command}")
        result = os.system(command)
        if result == 0:
            print(f"Archivo {output_file} descargado exitosamente desde {registry}/{project}/{artifact}")
        else:
            print(f"Error al descargar el archivo desde {registry}/{project}/{artifact}")
            raise Exception("Error al realizar el pull del archivo OCI")

class trafficData(HttpUser):
    wait_time = between(0.2, 0.9)
    reader = ReadFile()

    # Configuración del registro OCI
    registry = "registry.home-k8s.lab"  # Cambia esto por tu dominio del registro
    project = "project"  # Cambia esto por el nombre de tu proyecto en Harbor
    artifact = "asignaciones:latest"  # Cambia esto por la etiqueta de tu artefacto
    output_file = "data.json"  # Nombre del archivo o directorio descargado

    # Descargar el archivo JSON desde el registro OCI
    try:
        reader.pullFromOCI(registry, project, artifact, output_file)
        reader.loadFile(output_file)
    except Exception as e:
        print(f"Error al preparar los datos: {e}")
        # Detener Locust si no se puede descargar o cargar el archivo
        os._exit(1)

    @task
    def sendMessage(self):
        data = self.reader.getData()
        if data is not None:
            res = self.client.post("/insert", json=data)
            try:
                response = res.json()
                print(response)
            except Exception as e:
                print(f"Error al procesar la respuesta: {e}")
        else:
            print("No hay más datos")
            self.stop(True)
```

## 5. Ejecutar la prueba de carga con Locust
```bash
python -m venv venv
source venv/bin/activate
pip install locust
locust -f tu_script.py
http://0.0.0.0:8089
```
