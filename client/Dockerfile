# Usa una imagen base de Ubuntu para construir el binario
FROM ubuntu:latest AS builder

# Instala Go y otras dependencias necesarias
RUN apt-get update && \
    apt-get install -y wget git build-essential libgl1 libglx-mesa0 xorg-dev

# Descarga e instala Go
RUN wget https://golang.org/dl/go1.23.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz && \
    rm go1.23.3.linux-amd64.tar.gz

# Establece las variables de entorno para Go
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al directorio de trabajo
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila el binario
RUN go build -o client ./cmd

# Usa una imagen base de Ubuntu más pequeña para ejecutar el binario
FROM ubuntu:latest

# Instala las dependencias necesarias para ejecutar el binario
RUN apt-get update && \
    apt-get install -y libgl1 libglx-mesa0 libx11-6

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /root/

# Copia el binario desde la imagen de construcción
COPY --from=builder /app/client .
COPY startup.sh /root/startup.sh

# Establece la variable de entorno
ENV SERVER_URL=http://10.0.11.2:4000

# Comando para ejecutar el binario
RUN chmod +x /root/startup.sh

CMD ["sh", "-c", "/root/startup.sh && tail -f /dev/null"]