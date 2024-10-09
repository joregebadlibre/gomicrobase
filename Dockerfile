# Usa una imagen base oficial de Go
FROM golang:1.23.2-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app


# Copia el código fuente de la aplicación
COPY . .

# descarga las dependencias
RUN go mod tidy && go mod download

# Comando por defecto para ejecutar la aplicación
CMD ["go", "run", "/app/cmd/server/main.go"]