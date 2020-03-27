FROM golang:1.12.9 as build-stage
LABEL maintainer="Ren xiaohuo <lulu@zilliqa.com>"
WORKDIR /app
COPY ./ .
RUN go test -c -o ./test github.com/Zilliqa/gozilliqa-sdk/provider #gosetup

FROM alpine:latest
RUN CI=true go tool test2json -t ./test -test.v #gosetup



