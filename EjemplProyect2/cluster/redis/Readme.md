# Desplegar redis

## Paso 1: Crear el archivo de despliegue

[Despliegue de redis](redis.yaml)

## Paso 2: Desplegar redis

```bash
kubectl apply -f redis.yaml
```

## Paso 3: Verificar el despliegue

```bash
kubectl get pods
```

## Paso 4: Crear el servicio

[Servicio de redis](service.yaml)

```bash
kubectl apply -f service.yaml
```

## Paso 5: Verificar el servicio

```bash
kubectl get services
```

## Paso 6: Crear el hpa

[Horizontal Pod Autoscaler de redis](hpa.yaml)

```bash
kubectl apply -f hpa.yaml
```

## Paso 7 (Opcional): Verificar el acceso a redis

```bash
kubectl exec -it <<nombre del pod>> -- redis-cli
```