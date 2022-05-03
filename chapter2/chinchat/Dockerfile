FROM golang

ADD . /go/src/chinchat
WORKDIR /go/src/chinchat

RUN go install

ENTRYPOINT [ "/go/bin/chinchat" ]

EXPOSE 8080
