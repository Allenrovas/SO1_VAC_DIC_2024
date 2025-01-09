import json
import random

def generar_facultad():
    facultades = ["Ingenieria", "Medicina", "Derecho", "Arquitectura", "Economia", "Veterinaria", "Odontologia"]
    return random.choice(facultades)

def generar_curso():
    cursos = ["SO1", "LFP", "BD1", "SA", "AYD1", "SO2", "BD2", "AYD2", "SA", "IA1"]
    return random.choice(cursos)

def generar_carrera():
    carreras = ["Sistemas", "Industrial", "Civil", "Mecanica", "Electronica", "Quimica", "Biomedica"]
    return random.choice(carreras)

def generar_region():
    regiones = ["METROPOLITANA", "NORTE", "NORORIENTAL", "SURORIENTAL", "CENTRAL", "SUROCCIDENTAL", "NOROCCIDENTAL", "PETEN"]
    return random.choice(regiones)

def generar_json():
    data = {
        "curso": generar_curso(),
        "carrera": generar_carrera(),
        "facultad": generar_facultad(),
        "region": generar_region()
    }
    return data

def generar_archivo_json(cantidad,nombre_archivo):
    datos = []
    for _ in range(cantidad):
        datos.append(generar_json())
    
    with open(nombre_archivo, "w") as file:
        json.dump(datos, file, indent=4)


#Preguntar al usuario cuantos registros desea generar
cantidad = int(input("Ingrese la cantidad de registros a generar: "))

#Generar el archivo json
generar_archivo_json(cantidad, "data.json")