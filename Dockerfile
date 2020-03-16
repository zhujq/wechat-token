FROM alpine:latest

ADD account.json /account.json
ADD entrypoint.sh /entrypoint.sh
ADD wechat-token /wechat-token

RUN  chmod +x /wechat-token && chmod 777 /entrypoint.sh

ENTRYPOINT  /entrypoint.sh 

EXPOSE 1080
