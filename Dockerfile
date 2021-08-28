FROM golang:1.16

RUN apt update && \
  apt upgrade -y && \
  apt -y install git build-essential postgresql-client

ENV APP_HOME /go/src/app
COPY . $APP_HOME
WORKDIR $APP_HOME

RUN go mod download && \
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
