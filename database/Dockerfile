# Usa la imagen oficial de MySQL como base
FROM mysql:8.0

# Configura las variables de entorno para MySQL
ENV MYSQL_ROOT_PASSWORD=root_password
ENV MYSQL_DATABASE=distnetdb
ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=password

# Copia el archivo SQL al contenedor
COPY distnetdb.sql /docker-entrypoint-initdb.d/

# Exponer el puerto de MySQL
EXPOSE 3306