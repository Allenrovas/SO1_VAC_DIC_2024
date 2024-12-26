import json
import random

def generar_carnet():
    anio = random.choice(range(2020, 2024))
    numeros_aleratorios = random.randint(00000, 99999)
    carnet = f"{anio}{numeros_aleratorios}"
    return int(carnet)

def generar_nombre():
    nombres = ["Juan", "Pedro", "Maria", "Jose", "Carlos", "Ana", "Luis", "Luisa", "Sofia", "Andrea"]
    return random.choice(nombres)

def generar_curso():
    cursos = ["SO1", "LFP", "BD1", "SA", "AYD1"]
    return random.choice(cursos)

def generar_nota():
    return random.choice([50, 60, 70, 80, 90, 100])

def generar_semestre():
    return random.choice(['1S', '2S'])

def generar_json():
    data = {
        "carnet": generar_carnet(),
        "nombre": generar_nombre(),
        "curso": generar_curso(),
        "nota": generar_nota(),
        "semestre": generar_semestre(),
        "year": 2024
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