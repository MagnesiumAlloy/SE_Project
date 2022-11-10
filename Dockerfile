FROM golang:1.17

ENV GO111MODULE=on \
GOPROXY=https://goproxy.cn


WORKDIR /build

COPY . .

RUN go build main.go

EXPOSE 8080
ENTRYPOINT ["./main"]