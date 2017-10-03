FROM alpine:latest

LABEL maintainer="mark@amerine.net"

RUN mkdir -p /app
WORKDIR "/app"

ADD .docker_build/msgpack-dumper /app/msgpack-dumper

CMD ["/app/msgpack-dumper"]