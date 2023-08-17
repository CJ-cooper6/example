FROM golang:1.20.6-bullseye as build
WORKDIR /app
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
COPY . .
RUN go mod tidy
RUN go build -o compose_demo .

FROM ubuntu:latest
WORKDIR /usr/local/demo
COPY --from=build /app/compose_demo .
ENTRYPOINT ["./compose_demo"]