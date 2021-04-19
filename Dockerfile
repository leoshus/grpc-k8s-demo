FROM golang:1.16.3

COPY . /app/src/grpc-demo
WORKDIR /app/src/grpc-demo

RUN go get -d -v ./...
RUN go install -gcflags=all="-N -l " ./...


