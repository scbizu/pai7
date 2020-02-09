FROM golang:1.13 AS BUILDER

WORKDIR /project/pai7

ADD . /project/pai7

RUN export GO11MODULE="on" && go build -o pai7 .

FROM alpine:latest

RUN mkdir /app

COPY --from=BUILDER /project/pai7/pai7 /app/api7

ENTRYPOINT ["/app/pai7", "server"]
