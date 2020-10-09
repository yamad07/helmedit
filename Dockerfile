FROM golang:1.15-alpine

LABEL maintainer Yusuke Yamada <yamada@ispec.tech>

RUN apk --update add git less openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*
