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