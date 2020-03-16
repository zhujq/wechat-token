FROM alpine:latest

ADD account.json /account.json
ADD entrypoint.sh /entrypoint.sh
ADD wechat-token /wechat-token

RUN  chmod +x /wechat-token

ENTRYPOINT  /entrypoint.sh 

EXPOSE 1080
