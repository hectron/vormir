FROM golang:1.16

RUN apt update && apt -y install git build-essential

ENV APP_HOME /go/src/app
WORKDIR $APP_HOME

