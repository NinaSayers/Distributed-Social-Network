from alpine

run echo "net.ipv4.ip_forward=1" | tee -a /etc/sysctl.conf

cmd while true; do sleep 1; done