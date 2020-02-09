FROM golang:1.13-alpine

WORKDIR /project

ADD . .

ENV GOPROXY=goproxy.cn

RUN export GO11MODULE="on" && go build -o pai7

ENTRYPOINT [ "./pai7" ,"server"]
