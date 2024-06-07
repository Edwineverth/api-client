# API CLIENT
## Structure
```
/project-root
│
├── dao
│   └── client_dao.go
│
├── db
│   └── mongo.go
│
├── repository
│   └── client_repository.go
│
├── handler
│   └── client_handler.go
│
├── .env
├── main.go
└── go.mod
```

## Descripción

La API de `client` permite gestionar clientes en una base de datos MongoDB. Proporciona operaciones CRUD para crear, leer, actualizar y eliminar clientes.

## Requisitos

- Docker
- Docker Compose

## Instalación

1. Clona el repositorio:

   ```bash
   git clone https://github.com/Edwineverth/api-client.git

2. Asegurarse que las variables de entorno esten en un archivo .env

```bash
MONGO_URI=mongodb://mongo-client:27017
MONGO_DB=clientDB
SERVER_PORT=8080
```
## Uso

Construir y ejecutar la aplicación con Docker Compose
Construye y ejecuta los contenedores:

```bash
Copiar código
docker-compose up --build
La aplicación estará disponible en http://localhost:8080.
```
## Nota:
Al correr la aplicación correrá un pequeño script para llenar con información inicial la DB.
En baso de querer insertar más información modificar el archivo.
```makefile
init-portfolio.js
```

## Licencia
Este proyecto está licenciado bajo la Licencia MIT.