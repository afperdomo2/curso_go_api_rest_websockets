
# API REST y Websockets en GO

Este proyecto es un curso completo para desarrollar una API REST y Websockets usando Go (Golang). Incluye autenticación JWT, manejo de rutas HTTP y comunicación en tiempo real.

## 📋 Características

- ✅ API REST completa con operaciones CRUD
- ✅ Autenticación y autorización con JWT
- ✅ Comunicación en tiempo real con Websockets
- ✅ Manejo de variables de entorno
- ✅ Routing avanzado con Gorilla Mux
- ✅ Estructura de proyecto escalable

## 🚀 Tecnologías Utilizadas

- **Go (Golang)**: Lenguaje de programación principal
- **Gorilla Mux**: Router HTTP para manejo de rutas
- **Gorilla Websocket**: Implementación de Websockets
- **JWT-Go**: Manejo de JSON Web Tokens para autenticación
- **GoDotEnv**: Carga de variables de entorno desde archivo .env

## 📦 Instalación

### 1. Inicializar el proyecto

```sh
# Inicializar el módulo de Go
go mod init afperdomo2/go/rest-ws
```

### 2. Instalar dependencias

```sh
# Instalar todas las dependencias necesarias
go get github.com/golang-jwt/jwt
go get github.com/gorilla/mux
go get github.com/gorilla/websocket
go get github.com/joho/godotenv
```

### 3. Configurar variables de entorno

Crear un archivo `.env` en la raíz del proyecto:

```env
PORT=5050
JWT_SECRET=tu_clave_secreta_jwt
DATABASE_URL=tu_url_de_base_de_datos
```

## 🔧 Uso

### Ejecutar el servidor

```sh
go run main.go
```

El servidor se ejecutará en `http://localhost:5050` (o el puerto configurado en `.env`)

## 🛠️ Desarrollo

### Prerrequisitos

- Go 1.19 o superior
- Git

### Comandos útiles

```sh
# Ejecutar el proyecto
go run main.go

# Compilar el proyecto
go build -o api-server

# Ejecutar tests
go test ./...

# Formatear código
go fmt ./...

# Verificar dependencias
go mod tidy
```

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver `LICENSE` para más detalles.

## 👨‍💻 Autor

**Andrés Perdomo** - [@afperdomo2](https://github.com/afperdomo2)
