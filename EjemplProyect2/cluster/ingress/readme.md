# Pasos para realizar el ingress

## Paso 1: Instalar nginx

```bash
kubectl create ns nginx-ingress
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install nginx-ingress ingress-nginx/ingress-nginx -n nginx-ingress
kubectl get services -n nginx-ingress
```

## Paso 2: Crear el archivo de despliegue de ingress

Sustituir `<TU_IP>` por la IP del servidor donde se encuentra el cluster de Kubernetes.

Al realizar el `kubectl get services -n nginx-ingress` se obtiene la IP del servicio `ingress-nginx-controller`.

Se debe copiar la EXTERNAL-IP del LoadBalancer y sustituir `<TU_IP>` por esta IP.

Y se debe sustituir el `<NOMBRE_DEL_SERVICIO>` por el nombre del servicio que se
desea exponer. En este caso sera la API de GRPC.

[Despliegue de ingress](ingress.yaml)

```bash
kubectl apply -f ingress.yaml
```