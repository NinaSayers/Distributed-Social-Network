#!/bin/bash

# Verifica si docker y docker-compose están instalados
if ! command -v docker &> /dev/null
then
    echo "Docker no está instalado. Por favor, instálalo primero."
    exit 1
fi

if ! command -v docker-compose &> /dev/null
then
    echo "docker-compose no está instalado. Por favor, instálalo primero."
    exit 1
fi

# Construir las imágenes y levantar los contenedores
echo "Construyendo imágenes y levantando contenedores..."
docker-compose up -d --build

# Verificar el estado de los contenedores
echo "Verificando el estado de los contenedores..."
docker-compose ps

# Mensajes de éxito y accesos
echo "Sistema levantado con éxito!"
echo "Accede a los contenedores desde:"
echo "  Cliente1: http://10.0.10.2:3000"
echo "  Servidor: http://10.0.11.2:4000"
echo "  Base de Datos: http://10.0.11.3:3306"