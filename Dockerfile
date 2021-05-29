FROM golang as builder

RUN mkdir -p /src/notification
WORKDIR      /src/notification
ADD .  .
RUN CGO_ENABLED=0 GOOS=linux go build -o /notification-server

FROM alpine:latest
MAINTAINER "L. Jiang <l.jiang.1024@gmail.com>"
COPY --from=builder /notification-server /
RUN apk add --no-cache tzdata
RUN chmod +x /notification-server
EXPOSE 80
CMD ["/notification-server", "-c", "/server.toml"]