FROM golang:1.9

WORKDIR /go/src/github.com/guanw/kubernetes-service-discovery/client
COPY . .
RUN make build-linux
EXPOSE 8081
CMD ["/go/src/github.com/guanw/kubernetes-service-discovery/client/client_binary_unix"]