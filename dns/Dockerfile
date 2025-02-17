FROM docker.uclv.cu/golang:1.20-alpine

WORKDIR /dns

COPY ./ ./

EXPOSE 53/udp

CMD ["go", "run", "main.go"]