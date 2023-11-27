#!/usr/bin/env bash
docker stop labs
docker rm labs
docker image prune
docker build --rm -t labs:alpha .
docker run -d -p 8082:8080 --name labs labs:alpha