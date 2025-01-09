# Usamos una imagen de Rust con alpine como base
FROM rust:latest AS builder

# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos los archivos de nuestro proyecto
COPY . .

# Compilamos el proyecto
RUN cargo build --release

#Exponemos el puerto 8080
EXPOSE 8080

# Ejecutamos el binario
CMD ["./target/release/redis_storage"]