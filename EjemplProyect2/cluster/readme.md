# Notas

## Secret Harbor

En el archivo `secret-harbor.yaml` se encuentran las credenciales para acceder a Harbor. Estas credenciales son necesarias para poder acceder a Harbor desde el cluster de Kubernetes.

Se debe reemplaza <TU_CREDENCIAL_BASE64> con las credenciales de Harbor codificadas en base64. Puedes generarlas con:

```bash
cat ~/.docker/config.json | base64
```

