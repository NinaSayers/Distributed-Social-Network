#!/bin/sh

# Habilitar el reenvío de paquetes IP
echo "Habilitando el reenvío de paquetes IP..."
sysctl -w net.ipv4.ip_forward=1

# Configurar las interfaces de red
echo "Configurando interfaz eth0 (red test_kademlia)..."
ip addr add 10.0.10.254/24 dev eth0
ip link set eth0 up

echo "Configurando interfaz eth1 (red client)..."
ip addr add 10.0.11.254/24 dev eth1
ip link set eth1 up

# Configurar iptables para permitir el tráfico entre las redes
echo "Configurando iptables..."
iptables -A FORWARD -i eth0 -o eth1 -j ACCEPT
iptables -A FORWARD -i eth1 -o eth0 -j ACCEPT

# Habilitar NAT solo en la interfaz que tiene acceso a Internet (eth0)
echo "Configurando NAT en eth0..."
iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE

# Mantener el contenedor en ejecución
echo "Router configurado. Manteniendo el contenedor en ejecución..."
tail -f /dev/null