FROM alpine

# Instala las herramientas necesarias
RUN apk add --no-cache \
    iproute2 \
    tcpdump \
    iptables

# iproute2 --> para comandos como ip y tc
# tcpdump --> para depurar el trafico de red
# iptables --> para configurar reglas de firewall

# Configura el reenvío de paquetes IP
RUN echo "net.ipv4.ip_forward=1" >> /etc/sysctl.conf

# Copia un script de inicio al contenedor
COPY start-router.sh /usr/local/bin/start-router.sh
RUN chmod +x /usr/local/bin/start-router.sh

# Establece el script de inicio como comando por defecto
CMD ["/usr/local/bin/start-router.sh"]