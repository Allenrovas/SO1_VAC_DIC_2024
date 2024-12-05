#Usa una imagen de node.js
FROM node:22-alpine

#Establecer el directorio de trabajo dentro del contenedor
WORKDIR /usr/src/app
COPY package*.json ./

#Copiar el resto de los archivos
COPY . .

#Construye la aplicación para producción
RUN npm run build

#Exponer el puerto 3000
EXPOSE 3000

#Ejecutar la aplicación
CMD ["npm", "start"]