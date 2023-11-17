FROM golang:latest
WORKDIR /

ADD /tcpclient /tcpclient
ADD /start.sh /start.sh

WORKDIR /tcpclient/internal

RUN go build main.go
RUN mv main ../../main

RUN apt-get update -y
RUN apt-get install -y iputils-ping

WORKDIR /

CMD ./start.sh