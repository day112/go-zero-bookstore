#!/bin/sh

#     --url 'http://localhost:8888/check?book=etcd'\

for (( i = 0; i < 10; i++ )); do

echo ${i}
curl --request GET -sL \
     --url 'http://192.168.31.132:31443/check?book=go'\

echo /n
sleep 0.1

done


