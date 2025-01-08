# URL Shortener

Este proyecto implementa un servicio de acortamiento de URLs escrito en Go. Permite a los usuarios enviar URLs largas y recibir una versión corta que redirige a la URL original. Además, proporciona información estadística sobre el uso de las URLs acortadas.

## Características

- **Acortamiento de URLs**: Convierte URLs largas en códigos cortos.
- **Redirección**: Permite redirigir desde la URL corta a la original.
- **Estadísticas**: Ofrece datos estadísticos sobre las URLs acortadas.

## Requisitos

- Go 1.20+

## Instalación

1. Clona el repositorio:
   ```bash
   git clone <url-del-repositorio>
   cd <nombre-del-repositorio>
   ```
2. Instala las dependencias:
   ```bash
   go mod tidy
   ```
3. Ejecuta el servidor:
   ```bash
   go run main.go
   ```

## Uso

El servicio corre por defecto en `http://localhost:8080`.

### 1. Acortar una URL

**Endpoint**: `/shorten`

**Método**: `POST`

**Cuerpo de la solicitud** (JSON):

```json
{
  "url": "https://example.com"
}
```

**Respuesta** (JSON):

```json
{
  "short_url": "http://localhost:8080/{shortCode}"
}
```

### 2. Redirigir a la URL original

**Endpoint**: `/{shortCode}`

**Método**: `GET`

El navegador será redirigido a la URL original asociada al código corto.

### 3. Obtener estadísticas de una URL corta

**Endpoint**: `/stats/{shortCode}`

**Método**: `GET`

**Respuesta** (JSON):

```json
{
  "original_url": "https://example.com",
  "clicks": 42
}
```

## Arquitectura

El proyecto utiliza las siguientes bibliotecas y componentes:

- **mux**: Framework para manejar rutas HTTP.
- **sync.RWMutex**: Asegura acceso concurrente seguro al mapa de URLs.
- **crypto/rand**: Genera códigos únicos para las URLs cortas.

## Contribución

1. Haz un fork del proyecto.
2. Crea una rama para tu función (`git checkout -b feature/nueva-funcionalidad`).
3. Realiza los cambios necesarios y haz commit (`git commit -m "Agrega nueva funcionalidad"`).
4. Envía un pull request.

## Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo `LICENSE` para más información.

