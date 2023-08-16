FROM golang:1.20.6 as build
LABEL authors="cooper"
WORKDIR /app
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
COPY . .
RUN go build -o demo .

FROM ubuntu:latest
WORKDIR /usr/local/demo

COPY --from=build /app/demo .

ENTRYPOINT ["./demo"]