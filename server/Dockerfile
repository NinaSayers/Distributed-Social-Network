# Usa una imagen base de Go para construir el binario
FROM golang:1.23-alpine AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al directorio de trabajo
COPY . .

# Compila el binario
RUN go mod download
RUN go build -o server ./cmd/api

# Usa una imagen base más pequeña para ejecutar el binario
FROM alpine:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /root/

# Copia el binario desde la imagen de construcción
COPY --from=builder /app/server .
COPY startup.sh /root/startup.sh

RUN chmod +x /root/startup.sh

# Expone el puerto en el que el servidor escuchará
EXPOSE 4000

# Comando para ejecutar el binario
CMD ["sh", "-c","/root/startup.sh && ./server"]