#!/bin/bash

# Script to run the Game of Life Docker container on Raspberry Pi with display support

# Allow X11 connections from Docker
xhost +local:docker

# Run the container with X11 display access
docker run --rm \
  -e DISPLAY=$DISPLAY \
  -v /tmp/.X11-unix:/tmp/.X11-unix:rw \
  --device /dev/dri \
  --name go-game-of-life \
  sondreevik/go-arm-app:latest

# Revoke X11 access after container exits
xhost -local:docker
