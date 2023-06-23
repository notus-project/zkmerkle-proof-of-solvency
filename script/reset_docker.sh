#!/usr/bin/env bash

docker ps -aq | xargs docker stop | xargs docker rm

docker volume prune -f
docker network prune -f
docker container prune -f


# sudo rm -rf $HOME/docker_data
# mkdir -p $HOME/docker_data/mysql_data
# mkdir -p $HOME/docker_data/kvrocks_data


docker run -d --name zkpos-redis -p 6379:6379 redis

# docker run -d --name zkpos-mysql -p 3306:3306  -v $HOME/docker_data/mysql_data:/var/lib/mysql -e MYSQL_USER=zkpos -e MYSQL_PASSWORD=zkpos@123 -e MYSQL_DATABASE=zkpos  -e MYSQL_ROOT_PASSWORD=zkpos@123 mysql
docker run -d --name zkpos-mysql -p 3306:3306 -e MYSQL_USER=zkpos -e MYSQL_PASSWORD=zkpos@123 -e MYSQL_DATABASE=zkpos  -e MYSQL_ROOT_PASSWORD=zkpos@123 mysql

# docker run -d --name zkpos-kvrocks -p 6666:6666 -v $HOME/docker_data/kvrocks_data:/var/lib/kvrocks apache/kvrocks
docker run -d --name zkpos-kvrocks -p 6666:6666 apache/kvrocks


docker ps -a