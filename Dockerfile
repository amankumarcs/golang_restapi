FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./. /go/src/github.com/user/myProject/app
WORKDIR /go/src/github.com/user/myProject/app

RUN go build .
	
EXPOSE 8080