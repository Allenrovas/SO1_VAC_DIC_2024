# Desplegar mongo

## Paso 1: Crear secrets y desplegar

Crear el archivo de despliegue de secrets

[Despliegue de secrets](secret.yaml)

```bash
kubectl apply -f secret.yaml
```

## Paso 2: Crear pvc y desplegar

Crear el archivo de despliegue de pvc

[Despliegue de pvc](pvc.yaml)

```bash
kubectl apply -f pvc.yaml
```


## Paso 3: Crear el archivo de despliegue y desplegar mongo


[Despliegue de mongo](mongo.yaml)

```bash
kubectl apply -f mongo.yaml 
```

## Paso 4: Verificar el despliegue

```bash
kubectl get pods
```

## Paso 5: Crear el servicio

[Servicio de mongo](service.yaml)

```bash
kubectl apply -f service.yaml
```

## Paso 6: Verificar el servicio

```bash
kubectl get services
```

## Paso 7: Crear el hpa

[Horizontal Pod Autoscaler de mongo](hpa.yaml)

```bash
kubectl apply -f hpa.yaml
```

## Paso 8 (Opcional): Verificar el acceso a mongo

```bash
kubectl exec -it <<nombre del pod>> -- mongo
```


