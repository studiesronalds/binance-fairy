#!/bin/bash
set -o allexport; source .env; set +o allexport
DATE=$(echo date)
export $DATE
echo $DOCKER_REPOSITORY_URL
make "$@"