# Usa una imagen base de Go para construir el binario
FROM golang:1.23-alpine AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al directorio de trabajo
COPY . .

# Compila el binario
RUN go mod download
RUN go build -o client ./cmd

# Usa una imagen base más pequeña para ejecutar el binario
FROM alpine:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /root/

# Copia el binario desde la imagen de construcción
COPY --from=builder /app/client .

# Establece la variable de entorno
ENV SERVER_URL=http://app:4000

# Comando para ejecutar el binario
CMD ["./client"]