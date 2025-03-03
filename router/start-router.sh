#!/bin/sh

# Espera a que las interfaces de red estén disponibles (opcional, pero útil)
sleep 5

# Configura las interfaces de red (ADAPTAR eth0/eth1 SI ES NECESARIO)
# ip addr add 10.0.11.0/24 dev eth0
# ip link set eth0 up
echo "Configurando interfaz eth1 (red client)..."
ip addr add 10.0.11.1/24 dev eth1
ip link set eth1 up

# Habilita el reenvío de paquetes IP
sysctl -p

# Configura NAT#!/bin/sh

# Habilitar el reenvío de paquetes IP
echo "Habilitando el reenvío de paquetes IP..."
sysctl -w net.ipv4.ip_forward=1

# Configurar las rutas (opcional, Docker ya lo hace automáticamente)
echo "Configurando rutas..."
ip route add 10.0.10.0/24 dev eth0
ip route add 10.0.11.0/24 dev eth1

# Configurar iptables para permitir el tráfico entre las redes
echo "Configurando iptables..."
iptables -A FORWARD -i eth0 -o eth1 -j ACCEPT
iptables -A FORWARD -i eth1 -o eth0 -j ACCEPT

# Habilitar NAT (opcional, si es necesario)
iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
iptables -t nat -A POSTROUTING -o eth1 -j MASQUERADE

# Mantener el contenedor en ejecución
echo "Router configurado. Manteniendo el contenedor en ejecución..."
tail -f /dev/null
