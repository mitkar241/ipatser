#!/bin/env bash

docker-compose down
docker system prune -f
docker rmi ipatser_app
docker-compose up
