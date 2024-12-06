#Usar una imagen base de Python con Alpine
FROM python:3.11-alpine

#Establecer el directorio de trabajo
WORKDIR /app

#Copiar el archivo requirements.txt al directorio de trabajo
COPY requirements.txt requirements.txt
COPY app.py app.py

#Instalar las dependencias
RUN apk add --no-cache gcc musl-dev libffi-dev && \
    pip install --no-cache-dir -r requirements.txt

#Exponer el puerto 8000
EXPOSE 8000

#Ejecutar la aplicación
CMD ["python", "app.py"]