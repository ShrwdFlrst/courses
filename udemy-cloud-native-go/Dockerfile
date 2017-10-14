FROM golang:1.8.4-alpine3.6
MAINTAINER shrwdflrst

ENV SOURCES /go/src/github.com/shrwdflrst/cloud-native-go

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install -a

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT cloud-native-go
