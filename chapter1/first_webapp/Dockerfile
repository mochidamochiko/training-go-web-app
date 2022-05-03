FROM golang

ADD . /go/src/first_webapp
WORKDIR /go/src/first_webapp

RUN go install

ENTRYPOINT [ "/go/bin/first_webapp" ]

EXPOSE 8080
