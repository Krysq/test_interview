FROM golang:1.16
ADD . /go/main
WORKDIR /go/main
RUN cd src && go install
CMD cd src && go run main.go