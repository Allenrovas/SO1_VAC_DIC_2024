# Instalacion de Stress

## 1. Actualizacion de los paquetes

```sh
sudo apt-get update
```

## 2. Instalacion de Stress

```sh
sudo apt-get install -y stress 
```

## 3. Aplicar Stress

```sh
sudo stress --cpu 2 --timeout 60s
```
