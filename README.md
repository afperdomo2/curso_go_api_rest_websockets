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
go get golang.org/x/crypto/bcrypt
go get github.com/golang-jwt/jwt
go get github.com/gorilla/mux
go get github.com/gorilla/websocket
go get github.com/joho/godotenv
go get github.com/lib/pq
go get github.com/segmentio/ksuid
go get github.com/golang-jwt/jwt/v4
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
# Ejecutar el proyecto
go run main.go

# Ejecutar en modo watch (solo si se tiene instalado nodemon)
nodemon --exec "go run main.go" --ext go
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

## 🐳 Docker

```sh
# Levantar la base de datos para los usuarios
docker-compose up -d
```

## 🔎 Testear endpoints

**NOTA:** Los endpoints que tienen 🔒 son privados, se debe reemplazar el token, por uno vigente (generado en el Login)

### 🌎 Crear un nuevo usuario

```sh
curl --location 'http://localhost:5050/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "usuario123@gmail.com",
    "password": "contrasena123"
}'
```

### 🌎 Login

```sh
curl --location 'http://localhost:5050/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "usuario123@gmail.com",
    "password": "contrasena123"
}'
```

### 🔒 Consultar los datos del usuario logueado

```sh
curl --location 'http://localhost:5050/user-info' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3NTMwNzM5NDd9.1a8kPMPdMR-EZ_p7e0ZwPV-sr3wkzJa1Qp_8fmFFp4E'
```

### 🔒 Crear un Post

```sh
curl --location 'http://localhost:5050/posts' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3NTMwNzAyMzV9.EBjG2RFIFX7KTKhAuruW3qEPWMmSv8sK_X9FjqFjoyo' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Post nuevo",
    "content": "Contendio del post"
}'
```

### 🔒 Actualizar un Post existente

```sh
curl --location --request PUT 'http://localhost:5050/posts/1' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3NTMwNzAyMzV9.EBjG2RFIFX7KTKhAuruW3qEPWMmSv8sK_X9FjqFjoyo' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Nuevo título",
    "content": "Contendio actualizado"
}'
```

### 🔒 Borrar un Post existente

```sh
curl --location --request DELETE 'http://localhost:5050/posts/3' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3NTMwNzAyMzV9.EBjG2RFIFX7KTKhAuruW3qEPWMmSv8sK_X9FjqFjoyo'
```

### 🔒 Obtener un post por su ID

```sh
curl --location 'http://localhost:5050/posts/2' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwiZXhwIjoxNzUyOTg4MTU2fQ.QJEF2p18MeoALOxCAjQLKvz5xadIH9T-TC_ZaEvt2sY'
```

### 🔒 Listar todos los Posts (con paginación)

```sh
curl --location 'http://localhost:5050/posts?page=1&limit=20' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3NTMyNDI3MDd9.Vkni6WtKJz9-f1oFw9f9a3_d872IuoN6gbtGj_aUQ5o'
```

## 🌐 Conexión a WebSocket

El proyecto expone un endpoint WebSocket en `/ws` para comunicación en tiempo real. Puedes conectarte y enviar/recibir mensajes usando herramientas como `websocat`, `wscat` o desde el navegador.

### Ejemplo usando websocat

```sh
websocat ws://localhost:5050/ws
```

### Ejemplo usando curl (solo handshake)

```sh
curl --include --no-buffer ^
  --header "Connection: Upgrade" ^
  --header "Upgrade: websocket" ^
  --header "Host: localhost:5050" ^
  --header "Origin: http://localhost:5050" ^
  http://localhost:5050/ws
```

### Ejemplo usando JavaScript en el navegador

```js
const ws = new WebSocket('ws://localhost:5050/ws');
ws.onopen = () => {
  console.log('Conectado al WebSocket');
  ws.send(JSON.stringify({ type: 'mensaje', content: '¡Hola desde el cliente!' }));
};
ws.onmessage = (event) => {
  console.log('Mensaje recibido:', event.data);
};
```

Puedes enviar mensajes en formato JSON y recibir notificaciones en tiempo real. El servidor acepta múltiples clientes conectados simultáneamente.

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver `LICENSE` para más detalles.

## 👨‍💻 Autor

**Andrés Perdomo** - [@afperdomo2](https://github.com/afperdomo2)
