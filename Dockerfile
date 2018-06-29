FROM golang:1.10

ENV AWS_REGION us-west-2
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

ADD . $GOPATH/src/github.com/backendservice/todo

WORKDIR $GOPATH/src/github.com/backendservice/todo

RUN go get -u github.com/golang/dep/...
RUN dep ensure

EXPOSE 50080
RUN go build -o server *.go
CMD ["./server"]
