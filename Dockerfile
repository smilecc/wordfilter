FROM golang:1.10.3-stretch
EXPOSE 8091
ENV GOPATH /go/gopath
RUN go get github.com/huayuego/wordfilter; exit 0
WORKDIR /go/gopath/src/github.com/huayuego/wordfilter/service
ENTRYPOINT go run main.go 0.0.0.0:8091
