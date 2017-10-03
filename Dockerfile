FROM alpine:latest

MAINTAINER Mark Turner <mark@amerine.net>

WORKDIR "/opt"

ADD .docker_build/msgpack-dumper /opt/bin/msgpack-dumper

CMD ["/opt/bin/msgpack-dumper"]