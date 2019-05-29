FROM golang:alpine

WORKDIR /go/src/nicodive
COPY . .
ENV GO111MODULE=on

RUN apk --no-cache add git
RUN go build

EXPOSE 8080

ENTRYPOINT ["/go/src/nicodive/nicodive-api"]