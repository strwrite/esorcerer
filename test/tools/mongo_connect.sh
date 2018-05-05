#!/usr/bin/env bash

#docker run -it --network esorcerer-net \
#--rm mongo  sh -c 'exec mongo "some-mongo:27017/test" '


docker exec -it some-mongo mongo admin
