#!/bin/bash

# Crear redes
echo "Creando redes..."
docker network create --subnet=10.0.10.0/24 clients_net
docker network create --subnet=10.0.11.0/24 servers_net

# Construir imágenes
echo "Construyendo imágenes..."
docker build -t cliente:latest ./client
docker build -t database:latest ./database
docker build -t servidor:latest ./server
docker build -t my-router-image ./router

# Iniciar contenedor de la base de datos
echo "Iniciando contenedor de la base de datos..."
docker run -d --name database \
  --network servers_net --ip 10.0.11.100 \
  -e MYSQL_ROOT_PASSWORD=root_password \
  -e MYSQL_DATABASE=distnetdb \
  -e MYSQL_USER=user \
  -e MYSQL_PASSWORD=password \
  -p 3306:3306 \
  --cap-add NET_ADMIN \
  --restart unless-stopped \
  --privileged \
  mysql:8.0

# Iniciar contenedor del servidor
echo "Iniciando contenedor del servidor..."
docker run -d --name server_distnet \
  --network servers_net --ip 10.0.11.2 \
  -e DB_DSN=user:password@tcp(10.0.11.100:3306)/distnetdb?parseTime=true \
  -p 4000:4000 \
  --cap-add NET_ADMIN \
  --restart unless-stopped \
  --privileged \
  servidor:latest

# Iniciar contenedor del cliente
echo "Iniciando contenedor del cliente..."
docker run -d --name client1 \
  --network clients_net --ip 10.0.10.2 \
  -e SERVER_URL=http://10.0.11.2:4000 \
  -p 3000:3000 \
  --cap-add NET_ADMIN \
  --restart unless-stopped \
  --privileged \
  cliente:latest

# Iniciar contenedor del router
echo "Iniciando contenedor del router..."
docker run -d --name router \
  --network clients_net \
  --privileged \
  --cap-add NET_ADMIN \
  alpine:latest sh -c "sysctl -w net.ipv4.ip_forward=1 && tail -f /dev/null"

# Conectar el router a la red del servidor
docker network connect servers_net router


echo "Todos los contenedores están en funcionamiento."