FROM golang

RUN mkdir -p /go/src/market
WORKDIR /go/src/market

ADD . /go/src/market

RUN ./scripts/build.sh
