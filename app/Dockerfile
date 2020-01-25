FROM golang:1.11.5-alpine

RUN apk update && apk add vim git build-base

RUN go get -u github.com/gin-gonic/gin \
    && go get -u github.com/jinzhu/gorm \
    && go get -u github.com/jinzhu/gorm/dialects/mysql \
    && go get github.com/stretchr/testify \
    && go get github.com/oxequa/realize 

ARG AP_BASE_DIR=/go/src/app

WORKDIR /go/src/app