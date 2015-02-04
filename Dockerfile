FROM golang:1.4
ADD . /go/src/github.com/jmcarbo/consul-router
RUN go get github.com/jmcarbo/consul-router/...
RUN go install github.com/jmcarbo/consul-router
ENTRYPOINT /go/bin/consul-router
EXPOSE 8080
