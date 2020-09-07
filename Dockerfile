FROM golang:1.13 as hello-grpc

WORKDIR /workspace
ADD . /workspace

#RUN go get -d -v ./...

#build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -mod=vendor -o app /workspace/server.go && chmod +x app

CMD ["/workspace/app"]
EXPOSE 50051