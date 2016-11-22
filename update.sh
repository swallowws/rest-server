#!/bin/bash

git pull && \
go build rest-server.go && \
service rest-server stop && \
cp rest-server{,.toml} /opt/rest-server && \
service rest-server start && \