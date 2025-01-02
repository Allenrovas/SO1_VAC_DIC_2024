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