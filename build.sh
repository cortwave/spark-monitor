#!/bin/bash

trap 'rm ca-certificates.crt' EXIT
trap 'rm main' EXIT

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
curl -o ca-certificates.crt https://curl.haxx.se/ca/cacert.pem
docker build -t cortwave/spark-monitor:$1 .
