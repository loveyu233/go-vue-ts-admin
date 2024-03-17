FROM golang:1.19.13-alpine3.18 AS builder

WORKDIR /build

COPY . .

ENV GOPROXY=https://goproxy.cn,direct

RUN go mod download && go build -ldflags "-s -w" -o server main.go

FROM alpine:3.18

ENV TZ Asia/Shanghai

WORKDIR /app

COPY --from=builder /build/server /app/server
COPY --from=builder /build/config/config.yml /app/config/config.yml
COPY --from=builder /build/static/image/* /app/static/image/

EXPOSE 9666

CMD /app/server