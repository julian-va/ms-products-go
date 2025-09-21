# ms-products

Microservicio RESTful en Go para la gestión de productos.

## Descripción

Este proyecto implementa un microservicio que permite crear, consultar, actualizar y eliminar productos a través de una API REST. Está diseñado para ejecutarse localmente o en entornos de Kubernetes y utiliza almacenamiento en memoria para los productos.

## Características principales

- API REST para operaciones CRUD de productos.
- Almacenamiento en memoria.
- Middlewares para logging de peticiones.
- Configuración externa mediante `application.yml`.
- Preparado para despliegue en Kubernetes (archivos en la carpeta `k8`).
- Dockerfile para construcción de imagen.
- Inyección de dependencias.

## Estructura del proyecto

- `src/controller/`: Controladores de endpoints.
- `src/repository/`: Repositorios de datos.
- `src/service/`: Lógica de negocio.
- `src/router/`: Definición de rutas.
- `src/middleware/`: Middlewares HTTP.
- `src/config/`: Configuración e inyección de dependencias.
- `src/entity/`: Definición de entidades.
- `src/utils/`: Utilidades y helpers.
- `k8/`: Archivos de despliegue para Kubernetes.

## Endpoints principales

- `GET /products`: Listar productos.
- `POST /products`: Crear producto.
- `GET /products/{id}`: Consultar producto por ID.
- `PUT /products/{id}`: Actualizar producto.
- `DELETE /products/{id}`: Eliminar producto.
- `GET /ping`: Endpoint de salud.

## Ejecución local

1. Clona el repositorio.
2. Ejecuta `go run src/main.go`.

## Despliegue en Kubernetes

Utiliza los archivos en la carpeta `k8` para desplegar el microservicio en un clúster Kubernetes.

---

Desarrollado en Go. Listo para integración y despliegue en entornos modernos.
