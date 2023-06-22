#!/usr/bin/env bash

docker ps -aq | xargs docker stop | xargs docker rm

docker volume prune -f
docker network prune -f
docker container prune -f

sudo rm -rf /home/ubuntu/docker_data
mkdir -p /home/ubuntu/docker_data/mysql_data
mkdir -p /home/ubuntu/docker_data/kvrocks_data


docker run -d --name zkpos-redis -p 6379:6379 redis

docker run -d --name zkpos-mysql -p 3306:3306  -v /home/ubuntu/docker_data/mysql_data:/var/lib/mysql -e MYSQL_USER=zkpos -e MYSQL_PASSWORD=zkpos@123 -e MYSQL_DATABASE=zkpos  -e MYSQL_ROOT_PASSWORD=zkpos@123 mysql

docker run -d --name zkpos-kvrocks -p 6666:6666 -v /server/docker_data/kvrocks_data:/var/lib/kvrocks apache/kvrocks


docker ps -a