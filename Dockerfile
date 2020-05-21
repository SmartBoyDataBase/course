FROM golang:1.12-alpine as builder
RUN apk add git
COPY . /go/src/course
ENV GO111MODULE on
WORKDIR /go/src/course
RUN go get && go build

FROM alpine
MAINTAINER longfangsong@icloud.com
COPY --from=builder /go/src/course/course /
WORKDIR /
CMD ./course
ENV PORT 8000
EXPOSE 8000