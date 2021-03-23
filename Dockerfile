FROM golang:1.14-alpine as builder

RUN apk update && apk add git

RUN mkdir /app

ADD . /app

WORKDIR /app/src