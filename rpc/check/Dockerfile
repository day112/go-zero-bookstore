FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app
COPY ./check .
COPY ./etc .


FROM alpine

#RUN apk update --no-cache
#RUN apk add --no-cache ca-certificates
#RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai
ENV LOCAL_HOST 192.168.31.132

WORKDIR /app
COPY --from=builder app/check /app
COPY --from=builder app/check.yaml /app


CMD ["./check", "-f", "check.yaml"]
