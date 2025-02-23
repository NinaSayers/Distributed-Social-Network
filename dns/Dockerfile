FROM golang:1.24-alpine

WORKDIR /dns

COPY ./ ./

EXPOSE 53/udp

CMD ["go", "run", "main.go"]