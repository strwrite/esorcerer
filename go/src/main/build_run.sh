#!/usr/bin/env bash

docker build -t esorcerertest . && docker run --network esorcerer-net --rm -it --name esorcerer esorcerertest app