FROM golang:1.5.1
ENTRYPOINT ["/bin/http-hello-world"]

ADD . /go/src/github.com/sqdk/http-hello-world
RUN go get github.com/sqdk/http-hello-world

RUN go build -o /bin/http-hello-world github.com/sqdk/http-hello-world

ENTRYPOINT /bin/http-hello-world
