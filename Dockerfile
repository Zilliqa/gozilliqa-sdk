FROM golang:1.15.8-alpine
LABEL maintainer="Ren xiaohuo <lulu@zilliqa.com>"
WORKDIR /app
COPY ./ .
RUN apk add build-base
RUN go test -c -o ./test github.com/Zilliqa/gozilliqa-sdk/provider
RUN CI=true go tool test2json -t ./test -test.v