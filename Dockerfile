FROM golang

RUN mkdir -p /market
WORKDIR /market

ADD . /market

RUN ./scripts/build.sh
