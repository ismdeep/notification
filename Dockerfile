FROM golang:1.20 as builder
WORKDIR /src
COPY . .
RUN go build -o ./bin/notification-server -mod vendor -trimpath -ldflags '-s -w' github.com/ismdeep/notification/app/server

FROM debian:12
MAINTAINER "L. Jiang <l.jiang.1024@gmail.com>"
ENV TZ=Asia/Shanghai \
    SERVER_BIND=0.0.0.0:80
RUN set -e; \
    apt-get update; \
    apt-get upgrade -y; \
    apt-get install -y apt-transport-https ca-certificates tzdata
COPY --from=builder /src/bin/notification-server /usr/bin/
EXPOSE 80
CMD ["notification-server"]