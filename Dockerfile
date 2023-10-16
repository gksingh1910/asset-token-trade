FROM golang:latest
RUN mkdir /build
WORKDIR /build

ADD . /build/
RUN go build -o pcapp
ENTRYPOINT ["/build/pcapp"]