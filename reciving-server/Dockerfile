FROM golang:1.8-alpine
RUN mkdir -p /usr/src/app
COPY . /usr/src/app

WORKDIR /usr/src/app
RUN export GOPATH=$GOPATH:$PWD
RUN go build

EXPOSE 3333
CMD ["./app"]
