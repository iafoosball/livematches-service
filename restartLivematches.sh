#!/bin/bash
cd /home/joe/livematches-service
docker-compose kill -s SIGINT
docker stop $(docker ps -a | grep livematches | awk '{print $1}')
docker rm $(docker ps -a | grep livematches | awk '{print $1}')
docker-compose -f docker-compose.prod.yml up -d