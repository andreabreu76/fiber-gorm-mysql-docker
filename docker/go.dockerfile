FROM golang:1.18

RUN apt update && apt -y upgrade

WORKDIR /go/src/app/

EXPOSE 3000
