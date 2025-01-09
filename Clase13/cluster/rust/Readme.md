# Instalacion rust

## 1. Actualizar el sistema

```bash
sudo apt update && sudo apt upgrade -y
```

## 2. Instalar rust con rustup

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

## 3. Instrucciones del instalador

### 3.1. Presionar enter para instalar rust con la configuracion predeterminada

### 3.2 Esto instalara rust y cargo

## 4. Agregar rust al PATH

```bash
source $HOME/.cargo/env
```

## 5. Verificar la instalacion

```bash
rustc --version

cargo --version
```

## 6. Crear un nuevo proyecto

```bash
cargo new redis_storage
```
