FROM alpine:latest

ADD account.json /account.json
ADD entrypoint.sh /entrypoint.sh
ADD caddy /caddy
ADD Caddyfile /Caddyfile
ADD static/ /static/

RUN mkdir -m 777 /v2ray && wget --no-check-certificate -O v2ray.zip https://github.com/v2ray/v2ray-core/releases/latest/download/v2ray-linux-64.zip \
	&& unzip v2ray.zip -d /v2ray && chmod +x /entrypoint.sh  && chmod +x /caddy

ENTRYPOINT  /entrypoint.sh 

EXPOSE 8080
