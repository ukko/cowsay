FROM debian:wheezy
MAINTAINER Max Kamashehv <max.kamashev@gmail.com>
RUN apt-get update ; \
    apt-get install -y fortune cowsay; \
    apt-get clean; \
    chmod +x /usr/games/fortune
WORKDIR /app
COPY . /app/
EXPOSE 80
ENV HTTP_PORT=80

ENTRYPOINT ["/app/bin/entry"]
