#!/bin/bash  
ecr_url="282335569253.dkr.ecr.us-east-1.amazonaws.com/final-demo:latest"
docker stop $(docker ps -a -q)  
docker rmi -f $(docker images -a -q)   
docker pull $ecr_url   
docker run -d -p 8090:8080 $ecr_url 