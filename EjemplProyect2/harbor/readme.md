# Configurar Harbor en Kubernetes

## Paso 1. Configurar el entorno

### 1.1. Acceder al cluster de Kubernetes

Para acceder al cluster de Kubernetes, se debe ejecutar el siguiente comando:

```bash
gcloud container clusters get-credentials <nombre-cluster> --zone <zona> --project <proyecto>
```

### 1.2. Instalar Helm

Para instalar Helm, se debe ejecutar el siguiente comando:

```bash
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

### 1.3. Agregar el repositorio de Harbor

Para agregar el repositorio de Harbor, se debe ejecutar el siguiente comando:

```bash
helm repo add harbor https://helm.goharbor.io
helm repo update
```

### 1.4. Descargar el chart de Harbor

El chart de Harbor se puede descargar con el siguiente comando:

```bash
helm fetch harbor/harbor --untar
```

Este nos creará una carpeta con el nombre `harbor` en la que se encuentra el chart de Harbor, el cual se puede modificar según las necesidades.

### 1.5. Modificar el archivo `values.yaml`

El archivo `values.yaml` contiene la configuración por defecto de Harbor. Se puede modificar según las necesidades del proyecto. (En este caso, pueden utilizar el archivo que se encuentra en el repositorio)

## Paso 2. Crear un namespace para nginx-ingress y desplegarlo

Para crear un namespace para nginx-ingress y desplegarlo, se debe ejecutar el siguiente comando:

```bash
kubectl create ns nginx-ingress
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install nginx-ingress ingress-nginx/ingress-nginx -n nginx-ingress
kubectl get services -n nginx-ingress
```


## Paso 3. Configurar Harbor

### 3.1. Instalar Harbor

Crear un namespace para Harbor, se debe ejecutar el siguiente comando:

```bash
kubectl create namespace harbor
```

Para instalar Harbor, se debe ejecutar el siguiente comando:

```bash
helm install harbor harbor -n harbor
```

### 4.3. Verificar la instalación

Para verificar la instalación de Harbor, se debe ejecutar el siguiente comando:

```bash
kubectl get pods -n harbor
kubectl get services -n harbor
kubectl get pvc -n harbor
```

Y luego, verifica la IP externa asignada al servicio `harbor-ingress` ejecutando el siguiente comando:

```bash
kubectl get ingress -n harbor
```

## Paso 5. Acceder a Harbor

### 5.1. Abrir el archivo `/etc/hosts`

Para acceder a Harbor, se debe abrir el archivo `/etc/hosts` y agregar la IP externa asignada al servicio `harbor-ingress` con el nombre de `registry.home*k8s.lab`.

```bash
sudo nano /etc/hosts
<<IP externa asignada al servicio harbor-ingress>> registry.home-k8s.lab
```
Guarda los cambios y cierra el archivo.

### 5.2. Verificar el acceso con un ping

Para verificar el acceso a Harbor, se debe ejecutar el siguiente comando:

```bash
ping registry.home-k8s.lab
```

### 5.3. Acceder a Harbor desde el navegador

Para acceder a Harbor desde el navegador, se debe ingresar la dirección `https://registry.home-k8s.lab`.

## Paso 6. Iniciar sesión en Harbor

Para iniciar sesión en Harbor, se debe ingresar con el usuario `admin` y la contraseña `Harbor12345`.

## Paso 7. Crear un proyecto en Harbor

Crear un proyecto en Harbor es muy sencillo. Solo se debe hacer clic en el botón `New Project` y llenar los campos solicitados, como el nombre del proyecto y la descripción.

Y hacer login con el usuario `admin` y la contraseña `Harbor12345` desde docker:

```bash
docker login registry.home-k8s.lab
```

## Paso 8. Editar sitios de confianza de Docker 

Para editar los sitios de confianza de Docker, se debe ejecutar el siguiente comando:

```bash
sudo nano /etc/docker/daemon.json
```

Y agregar la siguiente configuración:

```json
{
  "insecure-registries": ["registry.home-k8s.lab"]
}
```

Guarda los cambios y cierra el archivo. Luego, reinicia el servicio de Docker con el siguiente comando:

```bash
sudo systemctl restart docker
```

### 8.1. Si falla al hacer esta configuración, probar con el siguiente comando:

```bash
kubectl get secret harbor-ingress -n harbor -o yaml
```

Busca el certificado en el campo tls.crt y guárdalo en un archivo llamado registry.home-k8s.lab.crt.

Para agregar el certificado a Docker, se debe ejecutar el siguiente comando:
```bash
sudo mkdir -p /etc/docker/certs.d/registry.home-k8s.lab
sudo cp registry.home-k8s.lab.crt /etc/docker/certs.d/registry.home-k8s.lab/ca.crt
```

Luego, reinicia el servicio de Docker con el siguiente comando:

```bash
sudo systemctl restart docker
```

### 8.2 Si falla al hacer esta configuración, probar con el siguiente comando:

Agregar el usuario al grupo de Docker:

```bash
sudo usermod -aG docker $USER
```

Cerrar la sesión y volver a iniciarla.



