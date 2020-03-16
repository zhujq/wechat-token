#cd /v2raybin
#echo -e "$CONFIG_JSON" > config.json
#if [ "$CERT_PEM" != "$KEY_PEM" ]; then
#  echo -e "$CERT_PEM" > cert.pem
#  echo -e "$KEY_PEM"  > key.pem
#fi


nohup ./caddy &
cd /v2ray
cp -f /config.json .
chmod +x v2ray v2ctl
./v2ray

