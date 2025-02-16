FROM alpine

# Instala las herramientas necesarias para sysctl
RUN apk add --no-cache procps

# Configura el reenvío de paquetes IP
RUN echo "net.ipv4.ip_forward=1" >> /etc/sysctl.conf

# Usa un script para ejecutar el comando sysctl y mantener el contenedor vivo
CMD sysctl -p && tail -f /dev/null