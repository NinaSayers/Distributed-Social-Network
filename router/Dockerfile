FROM alpine

# Instalar los paquetes necesarios, incluyendo iptables
RUN apk update && apk add --no-cache iproute2 iputils iptables

# Habilitar el reenvío de IP de forma persistente
RUN echo "net.ipv4.ip_forward=1" > /etc/sysctl.conf

# Agregar un script para habilitar el reenvío de IP y configurar iptables al iniciar el contenedor
CMD sysctl -p && /bin/sh