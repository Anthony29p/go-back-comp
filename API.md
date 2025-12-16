# Files API - Documentación

API para gestión de archivos en S3.

## Desarrollo

### Prerequisitos

Instala las siguientes herramientas:
- [Docker](https://www.docker.com/get-started)
- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html)
- [AWS CLI](https://aws.amazon.com/cli/)

### Configuración inicial

1. Instalar dependencias de Go:
```bash
make deps
```

2. Generar código de Wire:
```bash
make wire
```

### Desarrollo local

3. Iniciar servidor local:
```bash
make dev
```

La API estará disponible en: `http://localhost:3000`

### Deploy a AWS

```bash
make deploy
```

## Endpoints

### 1. Generar URL prefirmada para subir archivo

Genera una URL prefirmada que permite subir un archivo directamente a S3.

**Endpoint:** `POST /files/presigned/upload`

**Request:**
```json
{
  "fileName": "documento.pdf"
}
```

**Response (200):**
```json
{
  "fileId": "123e4567-e89b-12d3-a456-426614174000",
  "uploadUrl": "https://my-bucket.s3.amazonaws.com/..."
}
```

**Errores:**
- `400`: fileName es requerido
- `500`: Error al generar la URL prefirmada
