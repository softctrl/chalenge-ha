#!/bin/bash

cd && \
wget -c https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz && \
tar -zxvf go1.9.linux-amd64.tar.gz && PATH="${PATH}:~/go/bin" && \
mkdir Chalenge && cd Chalenge && cp -R /vagrant/* . && \
go test -v && go build -o Chalenge
