FROM golang:1.22-alpine

WORKDIR /server

RUN go get -u -v github.com/shadowsocks/go-shadowsocks2

EXPOSE ${port}

CMD ["go-shadowsocks2 -s", "'ss://AEAD_CHACHA20_POLY1305:${password}@:${port}'"]