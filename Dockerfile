FROM golang:1.16

RUN apt update && apt -y install git build-essential

ENV APP_HOME /go/src/app
COPY . $APP_HOME
WORKDIR $APP_HOME

RUN go mod download
