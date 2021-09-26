#!/usr/bin/env bash

docker build -t jeffssh/redirect:latest .
docker push jeffssh/redirect:latest
